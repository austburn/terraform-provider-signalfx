package signalfx

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var PaletteColors = map[string]int{
	"gray":       0,
	"blue":       1,
	"azure":      2,
	"navy":       3,
	"brown":      4,
	"orange":     5,
	"yellow":     6,
	"magenta":    7,
	"purple":     8,
	"pink":       9,
	"violet":     10,
	"lilac":      11,
	"iris":       12,
	"emerald":    13,
	"green":      14,
	"aquamarine": 15,
}

var FullPaletteColors = map[string]int{
	"gray":        0,
	"blue":        1,
	"azure":       2,
	"navy":        3,
	"brown":       4,
	"orange":      5,
	"yellow":      6,
	"magenta":     7,
	"purple":      8,
	"pink":        9,
	"violet":      10,
	"lilac":       11,
	"iris":        12,
	"emerald":     13,
	"green":       14,
	"aquamarine":  15,
	"red":         16,
	"gold":        17,
	"greenyellow": 18,
	"chartreuse":  19,
	"jade":        20,
}

func resourceAxisMigrateState(v int, is *terraform.InstanceState, meta interface{}) (*terraform.InstanceState, error) {
	switch v {
	case 0:
		return migrateAxisStateV0toV1(is)
	default:
		return is, fmt.Errorf("Unexpected schema version: %d", v)
	}
}

func migrateAxisStateV0toV1(is *terraform.InstanceState) (*terraform.InstanceState, error) {
	if is.Empty() || is.Attributes == nil {
		return is, nil
	}
	if v, ok := is.Attributes["max_value"]; ok {
		if f, err := strconv.ParseFloat(v, 32); err == nil && f == math.MaxFloat32 {
			delete(is.Attributes, "max_value")
		}
	}
	if v, ok := is.Attributes["min_value"]; ok {
		if f, err := strconv.ParseFloat(v, 32); err == nil && f == -math.MaxFloat32 {
			delete(is.Attributes, "min_value")
		}
	}
	if v, ok := is.Attributes["low_watermark"]; ok {
		if f, err := strconv.ParseFloat(v, 32); err == nil && f == -math.MaxFloat32 {
			delete(is.Attributes, "low_watermark")
		}
	}
	if v, ok := is.Attributes["high_watermark"]; ok {
		if f, err := strconv.ParseFloat(v, 32); err == nil && f == math.MaxFloat32 {
			delete(is.Attributes, "high_watermark")
		}
	}
	return is, nil
}

func timeChartResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the chart",
			},
			"program_text": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Signalflow program text for the chart. More info at \"https://developers.signalfx.com/docs/signalflow-overview\"",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the chart",
			},
			"unit_prefix": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "(Metric by default) Must be \"Metric\" or \"Binary\"",
			},
			"color_by": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "(Dimension by default) Must be \"Dimension\" or \"Metric\"",
			},
			"minimum_resolution": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The minimum resolution (in seconds) to use for computing the underlying program",
			},
			"max_delay": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "How long (in seconds) to wait for late datapoints",
				ValidateFunc: validateMaxDelayValue,
			},
			"disable_sampling": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "(false by default) If false, samples a subset of the output MTS, which improves UI performance",
			},
			"time_range": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ValidateFunc:  validateSignalfxRelativeTime,
				Description:   "From when to display data. SignalFx time syntax (e.g. -5m, -1h)",
				ConflictsWith: []string{"start_time", "end_time"},
			},
			"start_time": &schema.Schema{
				Type:          schema.TypeInt,
				Optional:      true,
				Description:   "Seconds since epoch to start the visualization",
				ConflictsWith: []string{"time_range"},
			},
			"end_time": &schema.Schema{
				Type:          schema.TypeInt,
				Optional:      true,
				Description:   "Seconds since epoch to end the visualization",
				ConflictsWith: []string{"time_range"},
			},
			"axis_right": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					SchemaVersion: 1,
					MigrateState:  resourceAxisMigrateState,
					Schema: map[string]*schema.Schema{
						"min_value": &schema.Schema{
							Type:        schema.TypeFloat,
							Optional:    true,
							Default:     -math.MaxFloat32,
							Description: "The minimum value for the right axis",
						},
						"max_value": &schema.Schema{
							Type:        schema.TypeFloat,
							Optional:    true,
							Default:     math.MaxFloat32,
							Description: "The maximum value for the right axis",
						},
						"label": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Label of the right axis",
						},
						"high_watermark": &schema.Schema{
							Type:        schema.TypeFloat,
							Optional:    true,
							Default:     math.MaxFloat32,
							Description: "A line to draw as a high watermark",
						},
						"high_watermark_label": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A label to attach to the high watermark line",
						},
						"low_watermark": &schema.Schema{
							Type:        schema.TypeFloat,
							Optional:    true,
							Default:     -math.MaxFloat32,
							Description: "A line to draw as a low watermark",
						},
						"low_watermark_label": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A label to attach to the low watermark line",
						},
						"watermarks": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": &schema.Schema{
										Type:        schema.TypeFloat,
										Required:    true,
										Description: "Axis value where the watermark line will be displayed",
									},
									"label": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Label to display associated with the watermark line",
									},
								},
							},
						},
					},
				},
			},
			"axis_left": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					SchemaVersion: 1,
					MigrateState:  resourceAxisMigrateState,
					Schema: map[string]*schema.Schema{
						"min_value": &schema.Schema{
							Type:        schema.TypeFloat,
							Optional:    true,
							Default:     -math.MaxFloat32,
							Description: "The minimum value for the left axis",
						},
						"max_value": &schema.Schema{
							Type:        schema.TypeFloat,
							Optional:    true,
							Default:     math.MaxFloat32,
							Description: "The maximum value for the left axis",
						},
						"label": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Label of the left axis",
						},
						"high_watermark": &schema.Schema{
							Type:        schema.TypeFloat,
							Optional:    true,
							Default:     math.MaxFloat32,
							Description: "A line to draw as a high watermark",
						},
						"high_watermark_label": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A label to attach to the high watermark line",
						},
						"low_watermark": &schema.Schema{
							Type:        schema.TypeFloat,
							Optional:    true,
							Default:     -math.MaxFloat32,
							Description: "A line to draw as a low watermark",
						},
						"low_watermark_label": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A label to attach to the low watermark line",
						},
						"watermarks": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": &schema.Schema{
										Type:        schema.TypeFloat,
										Required:    true,
										Description: "Axis value where the watermark line will be displayed",
									},
									"label": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Label to display associated with the watermark line",
									},
								},
							},
						},
					},
				},
			},
			"axes_precision": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Force a specific number of significant digits in the y-axis",
			},
			"axes_include_zero": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Force y-axes to always show zero",
			},
			"on_chart_legend_dimension": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Dimension to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: 'metric', 'plot_label' and any dimension.",
			},
			"legend_fields_to_hide": &schema.Schema{
				Type:          schema.TypeSet,
				Optional:      true,
				ConflictsWith: []string{"legend_options_fields"},
				Elem:          &schema.Schema{Type: schema.TypeString},
				Description:   "List of properties that shouldn't be displayed in the chart legend (i.e. dimension names)",
			},
			"legend_options_fields": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"property": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of a property to hide or show in the data table.",
						},
						"enabled": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "(true by default) Determines if this property is displayed in the data table.",
						},
					},
				},
				Optional:      true,
				ConflictsWith: []string{"legend_fields_to_hide"},
				Description:   "List of property and enabled flags to control the order and presence of datatable labels in a chart.",
			},
			"show_event_lines": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "(false by default) Whether vertical highlight lines should be drawn in the visualizations at times when events occurred",
			},
			"show_data_markers": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "(false by default) Show markers (circles) for each datapoint used to draw line or area charts",
			},
			"stacked": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "(false by default) Whether area and bar charts in the visualization should be stacked",
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Tags associated with the chart",
			},
			"plot_type": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "(LineChart by default) The default plot display style for the visualization. Must be \"LineChart\", \"AreaChart\", \"ColumnChart\", or \"Histogram\"",
				ValidateFunc: validatePlotTypeTimeChart,
			},
			"histogram_options": &schema.Schema{
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Options specific to Histogram charts",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"color_theme": &schema.Schema{
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "Base color theme to use for the graph.",
							ValidateFunc: validateFullPaletteColors,
						},
					},
				},
			},
			"viz_options": &schema.Schema{
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Plot-level customization options, associated with a publish statement",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"label": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The label used in the publish statement that displays the plot (metric time series data) you want to customize",
						},
						"color": &schema.Schema{
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "Color to use",
							ValidateFunc: validatePerSignalColor,
						},
						"axis": &schema.Schema{
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validateAxisTimeChart,
							Description:  "The Y-axis associated with values for this plot. Must be either \"right\" or \"left\"",
						},
						"plot_type": &schema.Schema{
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validatePlotTypeTimeChart,
							Description:  "(Chart plot_type by default) The visualization style to use. Must be \"LineChart\", \"AreaChart\", \"ColumnChart\", or \"Histogram\"",
						},
						"value_unit": &schema.Schema{
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validateUnitTimeChart,
							Description:  "A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes)",
						},
						"value_prefix": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "An arbitrary prefix to display with the value of this plot",
						},
						"value_suffix": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "An arbitrary suffix to display with the value of this plot",
						},
					},
				},
			},
			"synced": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether the resource in the provider and SignalFx are identical or not. Used internally for syncing.",
			},
			"last_updated": &schema.Schema{
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Latest timestamp the resource was updated",
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL of the chart",
			},
		},

		Create: timechartCreate,
		Read:   timechartRead,
		Update: timechartUpdate,
		Delete: timechartDelete,
	}
}

/*
  Use Resource object to construct json payload in order to create a time chart
*/
func getPayloadTimeChart(d *schema.ResourceData) ([]byte, error) {
	payload := map[string]interface{}{
		"name":        d.Get("name").(string),
		"description": d.Get("description").(string),
		"programText": d.Get("program_text").(string),
	}

	viz := getTimeChartOptions(d)
	if axesOptions := getAxesOptions(d); len(axesOptions) > 0 {
		viz["axes"] = axesOptions
	}
	// There are two ways to maniplate the legend. The first is keyed from
	// `legend_fields_to_hide`. Anything in this is marked as hidden. Unspecified
	// fields default to showing up in SFx's UI.
	if legendOptions := getLegendOptions(d); len(legendOptions) > 0 {
		viz["legendOptions"] = legendOptions
		// Alternatively, the `legend_options_fields` provides finer control,
		// allowing ordering and on/off toggles. This is preferred, but we keep
		// `legend_fields_to_hide` for convenience.
	} else if legendOptions := getLegendFieldOptions(d); len(legendOptions) > 0 {
		viz["legendOptions"] = legendOptions
	}

	if vizOptions := getPerSignalVizOptions(d); len(vizOptions) > 0 {
		viz["publishLabelOptions"] = vizOptions
	}
	if onChartLegendDim, ok := d.GetOk("on_chart_legend_dimension"); ok {
		if onChartLegendDim == "metric" {
			onChartLegendDim = "sf_originatingMetric"
		} else if onChartLegendDim == "plot_label" {
			onChartLegendDim = "sf_metric"
		}
		viz["onChartLegendOptions"] = map[string]interface{}{
			"showLegend":        true,
			"dimensionInLegend": onChartLegendDim,
		}
	}
	if len(viz) > 0 {
		payload["options"] = viz
	}
	if val, ok := d.GetOk("tags"); ok {
		tags := []string{}
		for _, tag := range val.([]interface{}) {
			tags = append(tags, tag.(string))
		}
		payload["tags"] = tags
	}

	return json.Marshal(payload)
}

func getPerSignalVizOptions(d *schema.ResourceData) []map[string]interface{} {
	viz := d.Get("viz_options").(*schema.Set).List()
	viz_list := make([]map[string]interface{}, len(viz))
	for i, v := range viz {
		v := v.(map[string]interface{})
		item := make(map[string]interface{})

		item["label"] = v["label"].(string)
		if val, ok := v["color"].(string); ok {
			if elem, ok := PaletteColors[val]; ok {
				item["paletteIndex"] = elem
			}
		}
		if val, ok := v["plot_type"].(string); ok && val != "" {
			item["plotType"] = val
		}
		if val, ok := v["axis"].(string); ok && val != "" {
			if val == "right" {
				item["yAxis"] = 1
			} else {
				item["yAxis"] = 0
			}
		}
		if val, ok := v["value_unit"].(string); ok && val != "" {
			item["valueUnit"] = val
		}
		if val, ok := v["value_suffix"].(string); ok && val != "" {
			item["valueSuffix"] = val
		}
		if val, ok := v["value_prefix"].(string); ok && val != "" {
			item["valuePrefix"] = val
		}

		viz_list[i] = item
	}
	return viz_list
}

func getAxesOptions(d *schema.ResourceData) []map[string]interface{} {
	axes_list_opts := make([]map[string]interface{}, 2)
	if tf_axis_opts, ok := d.GetOk("axis_right"); ok {
		tf_right_axis_opts := tf_axis_opts.(*schema.Set).List()[0]
		tf_opt := tf_right_axis_opts.(map[string]interface{})
		axes_list_opts[1] = getSingleAxisOptions(tf_opt)
	}
	if tf_axis_opts, ok := d.GetOk("axis_left"); ok {
		tf_left_axis_opts := tf_axis_opts.(*schema.Set).List()[0]
		tf_opt := tf_left_axis_opts.(map[string]interface{})
		axes_list_opts[0] = getSingleAxisOptions(tf_opt)
	}
	return axes_list_opts
}

func getSingleAxisOptions(axisOpt map[string]interface{}) map[string]interface{} {
	item := make(map[string]interface{})

	if val, ok := axisOpt["min_value"]; ok {
		if val.(float64) == -math.MaxFloat32 {
			item["min"] = nil
		} else {
			item["min"] = val.(float64)
		}
	}
	if val, ok := axisOpt["max_value"]; ok {
		if val.(float64) == math.MaxFloat32 {
			item["max"] = nil
		} else {
			item["max"] = val.(float64)
		}
	}
	if val, ok := axisOpt["label"]; ok {
		item["label"] = val.(string)
	}
	if val, ok := axisOpt["high_watermark"]; ok {
		if val.(float64) == math.MaxFloat32 {
			item["highWatermark"] = nil
		} else {
			item["highWatermark"] = val.(float64)
		}
	}
	if val, ok := axisOpt["high_watermark_label"]; ok {
		item["highWatermarkLabel"] = val.(string)
	}
	if val, ok := axisOpt["low_watermark"]; ok {
		if val.(float64) == -math.MaxFloat32 {
			item["lowWatermark"] = nil
		} else {
			item["lowWatermark"] = val.(float64)
		}
	}
	if val, ok := axisOpt["low_watermark_label"]; ok {
		item["lowWatermarkLabel"] = val.(string)
	}

	// special case: the axis object might exist, but it has no keys except
	// watermarks
	// in this case, we don't want to report an axis object to sfx at all
	if len(item) == 0 {
		return nil
	}
	return item
}

func getTimeChartOptions(d *schema.ResourceData) map[string]interface{} {
	viz := make(map[string]interface{})
	viz["type"] = "TimeSeriesChart"
	if val, ok := d.GetOk("unit_prefix"); ok {
		viz["unitPrefix"] = val.(string)
	}
	if val, ok := d.GetOk("color_by"); ok {
		viz["colorBy"] = val.(string)
	}
	if val, ok := d.GetOk("show_event_lines"); ok {
		viz["showEventLines"] = val.(bool)
	}
	viz["stacked"] = d.Get("stacked").(bool)
	if val, ok := d.GetOk("plot_type"); ok {
		viz["defaultPlotType"] = val.(string)
	}
	if val, ok := d.GetOk("axes_precision"); ok {
		viz["axisPrecision"] = val.(int)
	}
	if val, ok := d.GetOk("axes_include_zero"); ok {
		viz["includeZero"] = val.(bool)
	}

	programOptions := make(map[string]interface{})
	if val, ok := d.GetOk("minimum_resolution"); ok {
		programOptions["minimumResolution"] = val.(int) * 1000
	}
	if val, ok := d.GetOk("max_delay"); ok {
		programOptions["maxDelay"] = val.(int) * 1000
	}
	if val, ok := d.GetOk("disable_sampling"); ok {
		programOptions["disableSampling"] = val.(bool)
	}
	if len(programOptions) > 0 {
		viz["programOptions"] = programOptions
	}

	timeMap := make(map[string]interface{})
	if val, ok := d.GetOk("time_range"); ok {
		if ms, err := fromRangeToMilliSeconds(val.(string)); err == nil {
			timeMap["range"] = ms
			timeMap["type"] = "relative"
		}
	}
	if val, ok := d.GetOk("start_time"); ok {
		timeMap["start"] = val.(int) * 1000
		timeMap["type"] = "absolute"
		if val, ok := d.GetOk("end_time"); ok {
			timeMap["end"] = val.(int) * 1000
		}
	}
	if len(timeMap) > 0 {
		viz["time"] = timeMap
	}

	dataMarkersOption := make(map[string]interface{})
	dataMarkersOption["showDataMarkers"] = d.Get("show_data_markers").(bool)
	if chartType, ok := d.GetOk("plot_type"); ok {
		chartType := chartType.(string)
		if chartType == "AreaChart" {
			viz["areaChartOptions"] = dataMarkersOption
		} else if chartType == "LineChart" {
			viz["lineChartOptions"] = dataMarkersOption
		} else if chartType == "Histogram" {
			histogramChartOptions := make(map[string]interface{})
			if histogram_options, ok := d.GetOk("histogram_options"); ok {
				hOptions := histogram_options.(*schema.Set).List()[0].(map[string]interface{})
				if color_theme, ok := hOptions["color_theme"].(string); ok {
					if elem, ok := FullPaletteColors[color_theme]; ok {
						histogramChartOptions["colorThemeIndex"] = elem
						viz["histogramChartOptions"] = histogramChartOptions
					}
				}
			}
		}
	} else {
		viz["lineChartOptions"] = dataMarkersOption
	}

	return viz
}

func timechartCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*signalfxConfig)
	payload, err := getPayloadTimeChart(d)
	if err != nil {
		return fmt.Errorf("Failed creating json payload: %s", err.Error())
	}

	url, err := buildURL(config.APIURL, CHART_API_PATH, map[string]string{})
	if err != nil {
		return fmt.Errorf("[SignalFx] Error constructing API URL: %s", err.Error())
	}

	err = resourceCreate(url, config.AuthToken, payload, d)
	if err != nil {
		return err
	}
	// Since things worked, set the URL and move on
	appURL, err := buildAppURL(config.CustomAppURL, CHART_APP_PATH+d.Id())
	if err != nil {
		return err
	}
	d.Set("url", appURL)
	return nil
}

func timechartRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*signalfxConfig)
	path := fmt.Sprintf("%s/%s", CHART_API_PATH, d.Id())
	url, err := buildURL(config.APIURL, path, map[string]string{})
	if err != nil {
		return fmt.Errorf("[SignalFx] Error constructing API URL: %s", err.Error())
	}

	return resourceRead(url, config.AuthToken, d)
}

func timechartUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*signalfxConfig)
	payload, err := getPayloadTimeChart(d)
	if err != nil {
		return fmt.Errorf("Failed creating json payload: %s", err.Error())
	}
	path := fmt.Sprintf("%s/%s", CHART_API_PATH, d.Id())
	url, err := buildURL(config.APIURL, path, map[string]string{})
	if err != nil {
		return fmt.Errorf("[SignalFx] Error constructing API URL: %s", err.Error())
	}

	return resourceUpdate(url, config.AuthToken, payload, d)
}

func timechartDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*signalfxConfig)
	path := fmt.Sprintf("%s/%s", CHART_API_PATH, d.Id())
	url, err := buildURL(config.APIURL, path, map[string]string{})
	if err != nil {
		return fmt.Errorf("[SignalFx] Error constructing API URL: %s", err.Error())
	}

	return resourceDelete(url, config.AuthToken, d)
}

/*
  Validates the plot_type field against a list of allowed words.
*/
func validatePlotTypeTimeChart(v interface{}, k string) (we []string, errors []error) {
	value := v.(string)
	if value != "LineChart" && value != "AreaChart" && value != "ColumnChart" && value != "Histogram" {
		errors = append(errors, fmt.Errorf("%s not allowed; Must be \"LineChart\", \"AreaChart\", \"ColumnChart\", or \"Histogram\"", value))
	}
	return
}

/*
  Validates the axis right or left.
*/
func validateAxisTimeChart(v interface{}, k string) (we []string, errors []error) {
	value := v.(string)
	if value != "right" && value != "left" {
		errors = append(errors, fmt.Errorf("%s not allowed; must be either right or left", value))
	}
	return
}

func validateUnitTimeChart(v interface{}, k string) (we []string, errors []error) {
	value := v.(string)
	allowedWords := []string{
		"Bit",
		"Kilobit",
		"Megabit",
		"Gigabit",
		"Terabit",
		"Petabit",
		"Exabit",
		"Zettabit",
		"Yottabit",
		"Byte",
		"Kibibyte",
		"Mebibyte",
		"Gigibyte",
		"Tebibyte",
		"Pebibyte",
		"Exbibyte",
		"Zebibyte",
		"Yobibyte",
		"Nanosecond",
		"Microsecond",
		"Millisecond",
		"Second",
		"Minute",
		"Hour",
		"Day",
		"Week",
	}
	for _, word := range allowedWords {
		if value == word {
			return
		}
	}
	errors = append(errors, fmt.Errorf("%s not allowed; must be one of: %s", value, strings.Join(allowedWords, ", ")))
	return
}