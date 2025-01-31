/*
 * Charts API
 *
 * An API for creating, retrieving, updating, and deleting charts
 *
 * API version: 2.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package chart

type Options struct {
	AreaChartOptions *AreaChartOptions `json:"areaChartOptions,omitempty"`
	// *These options are only used for time series charts.*<br> Axis options for the left and right side of a time series chart. The object in the first element specifies options for the left side of the chart, and the object in the second element corresponds to the right side of the chart. Other elements are ignored.
	Axes []*Axes `json:"axes,omitempty"`
	// Specifies the digits SignalFx displays for values plotted on the chart. Choose a value that is sensible for the data; that is, ensure that the variations in the data are discernible using the specified precision. For example, if the data values usually fluctuate between 100000 and 100010, using a precision of 3 would result in a value of 100000 for every data point. Setting the precision to 6 would distinguish between 100001 and 100002.<br>**Note:** SignalFx only uses this value for time series charts.
	AxisPrecision int32 `json:"axisPrecision,omitempty"`
	// Specifies how to apply a color scheme to the values in the chart. If you want to apply color in a text chart, use HTML within the `markdown` property. The values you can use depend on the type of chart you specify in `options.type` -- <br>   * \"Text\" -- The property is ignored.   * \"Heatmap\" -- The `Range` and `Scale` enumerated types are allowed; the default is `Range`.   * \"List\" -- The `Dimension`, `Metric`, and `Scale` enumerated types are allowed; the default is `Metric`.   * \"SingleValue\" -- The `Dimension`, `Metric`, and `Scale` enumerated types are allowed; the default is `Metric`.   * \"TimeSeriesChart\" -- The `Dimension` and `Metric` enumerated types are allowed; the default is `Dimension`.
	ColorBy    string                    `json:"colorBy,omitempty"`
	ColorRange *HeatmapColorRangeOptions `json:"colorRange,omitempty"`
	ColorScale *ColorScale               `json:"colorScale,omitempty"`
	// An array that contains the information about a single color range, including both the color to display for that range and the borders of the range. The array specifies the entire range displayed in the secondary visualization or heatmap chart. You don't need to insert the elements in a specific order; they're automatically ordered by value in the display. The lowest value becomes the left border of the secondary visualization, and the highest value becomes the right border of the secondary visualization.
	ColorScale2 []*SecondaryVisualization `json:"colorScale2,omitempty"`
	// Specifies the plot type to use for this chart --   * **LineChart** -- A plot with datapoints connected by a series of straight lines.   * **AreaChart** -- Similar to a LineChart, with the area between the plot line and the x-axis filled with the line color.   * **Column** -- Shaded vertical bars with their bottom on the x-axis and their top at the datapoint value. The bars aren't     connected.   * **Histogram** -- Colored rectangular bins, with the color representing the number of datapoints for a value. Depending on how you set up your histogram, a green bar can indicate a higher number of datapoints for a particular value than a red     bar.
	DefaultPlotType          string                      `json:"defaultPlotType,omitempty"`
	EventPublishLabelOptions []*EventPublishLabelOptions `json:"eventPublishLabelOptions,omitempty"`
	// Specifies an array of properties that should be grouped together in a histogram chart. Each array element must contain a key specified in `customProperties` or a valid default custom property. The first array element is a property that determines the top-level grouping; the property in the second element determines the grouping within each first-level grouping. All other elements in the array are ignored. The [Charts Overview](https://developers.signalfx.com/v2/reference.html#charts-overview-1) section shows  a heatmap that groups CPU utilization by AWS availability zone as the primary grouping and number of host CPU cores as the secondary grouping.<br> **Note** This options is only available if `options.type` is `HeatMap`.
	GroupBy               []string               `json:"groupBy,omitempty"`
	HistogramChartOptions *HistogramChartOptions `json:"histogramChartOptions,omitempty"`
	// Determines how 0 is treated when dynamically calculating Y-axis values. If `true` 0 is included in the calculation; otherwise, SignalFx ignores datapoints containinig 0. This is a chart-wide setting applied to all the axes. However, if `options.axes.min` or `options.axes.max` exclude 0 from the displayed range for an axis, 0 isn't displayed on that axis even if `includeZero` is `true`.<br> **Note** Only available if the `options.type` is `TimeSeriesChart`
	IncludeZero      bool              `json:"includeZero,omitempty"`
	LegendOptions    *DataTableOptions `json:"legendOptions,omitempty"`
	LineChartOptions *LineChartOptions `json:"lineChartOptions,omitempty"`
	// The contents of a text chart, using GitHub-Flavored Markdown (**GFM**) or HTML.<br> **Note** Only available if `options.type` is set to `Text`.
	Markdown string `json:"markdown,omitempty"`
	// Indicates the number of significant digits to the right of the decimal point to use for plotted datapoints. The integer part of the datapoint is used, regardless of the specified precision. Choose a value so that variations in the data are visible for the specified precision. For example, if data values are in the range 0.001 to 0.01, set `maximumPrecision` to 4. If you don't specify a value, SignalFx adjusts the precision to fit the available display area.<br> **Note** Only available if `options.type` is `List` or `SingleValue`.
	MaximumPrecision     int32           `json:"maximumPrecision,omitempty"`
	OnChartLegendOptions *LegendOptions  `json:"onChartLegendOptions,omitempty"`
	ProgramOptions       *GeneralOptions `json:"programOptions,omitempty"`
	// An array of objects that contain general options for a specific SignalFlow statement applied to the chart. The order of elements in the list is retained, but has no effect. SignalFx uses label values to  map options to specific `publish() statements` in the SignalFlow statement. Heatmaps do not support full `publishLabelOptions` elements; the description of each property shows its supported chart types.<br> **Note** Available if `options.type` is one of `TimeSeriesChart`, `List`, `SingleValue`, or `Heatmap`.
	PublishLabelOptions []*PublishLabelOptions `json:"publishLabelOptions,omitempty"`
	// Specifies how often, in milliseconds, SignalFx should refresh the chart with data. If SignalFx can detect the resolution of the data, the default is 10,000 (10 seconds); otherwise, the default is 3,600,000 (1 hour).<br> **Note** This option is only available if `options.type` is `List` or `SingleValue`.
	RefreshInterval int32 `json:"refreshInterval,omitempty"`
	// Specifies the secondary visualization to use with `List` or `SingleValue` charts. See the [Charts Overview](https://developers.signalfx.com/reference#charts-overview-1) section for examples of the available visualizations.<br> **Note** This option is only available if `option.type` is `List` or `SingleValue`.
	SecondaryVisualization string `json:"secondaryVisualization,omitempty"`
	// If `true`, SignalFx displays a vertical line on the chart at the point that an event occurs.<br> **Note** This option is only available if `options.type` is `TimeSeriesChart`.
	ShowEventLines bool `json:"showEventLines,omitempty"`
	// If `true`, SignalFx displays a sparkline for the value shown in a SingleValue chart. For example, the following chart contains a sparkline<br> ![SingleValue chart with sparkline]()<br> **Notes**   * This option is only available if `options.type` is `SingleValue`.   * The value specified for `options.secondaryVisualization` overrides `options.showSparkLine`.
	ShowSparkLine bool `json:"showSparkLine,omitempty"`
	// Specifies how to sort the entries in a list chart. The first character of the value must be either   * A minus sign (`-`) sets descending order for the data.   * A plus sign (`+`) sets ascending order for the data. The rest of the value specifies the sort criteria --   * A keyword that represents a datapoint value, a metric, or the `publish()` label of the SignalFlow statement that     generates the data   * A dimension that's available to the cart. <br> The default is to sort on the `publish()` label in ascending order; this corresponds to the **Auto** option in the SignalFx web UI. <br> Values that don't map to a keyword or valid dimension are sorted with the default specification. <br> Data that contains `null` sorts to the start of the list for ascending order and to the end of the list for descending order.<br> **Note** This option is only available if `option.type` is `List`.
	SortBy string `json:"sortBy,omitempty"`
	// If set to `Ascending`, the chart displays values from low to high; otherwise, the chart displays values from high to low.<br> **Note** This option is only available if `options.type` is `HeatMap`.
	SortDirection string `json:"sortDirection,omitempty"`
	// The custom property to use as the sorting criteria for the heatmap chart. The value should map to a valid key name specified in the customProperties property or to a valid custom property available by default. If no value is specified, the values are displayed without an obvious order. The name of a custom property to use as the sort criteria for a heatmap chart. The value  must be a valid key specified by the `customProperty` property or a valid default custom property. If you don't specify a value, the chart displays the values without any obvious order.<br> **Note** This option is only available if `options.type` is `HeatMap`.
	SortProperty string `json:"sortProperty,omitempty"`
	// Controls the display of stacked plots in a TimeSeriesChart. If `true`, SignalFx stacks charts and adds datapoints together to create the values indicated by the axis labels. You can still see individual values when you highlight them within the colored area representing a single plot. You can also view individual values in the data table.<br> **Note** This option is only available if `options.type` is `TimeSeriesChart` **and** `options.defaultPlotType` is `AreaChart` or `ColumnChart`.
	Stacked bool                `json:"stacked,omitempty"`
	Time    *TimeDisplayOptions `json:"time,omitempty"`
	// Determines if the timestamp displays below displayed values. If `true`, the timestamp for a value appears below the value in the chart.<br> **Note** This option is only available if `options.type` is `Heatmap` or `SingleValue`.
	TimestampHidden bool `json:"timestampHidden,omitempty"`
	// The type of chart or visualization to use for the data\\:<br>   * **Heatmap** -- Displays values in an ordered set of colored boxes. The color indicates the \"health\" of the       value.   * **List** -- Display each value in its own row within a plot that shows recent changes   * **SingleValue** -- Displays a single value with a color that indicates the health of the value.   * **Text** -- Displays pre-defined text, optionally formatted with Markdown markup   * **TimeSeriesChart** -- Displays multiple data points, retaining historical data, in one of four modes\\:       * **Line** -- Displays datapoints as chart points connected by straight lines. Each set of chart points and           lines has a color specific to that set.       * **Area** -- Similar to a Line chart, but the area between a line and the           X-axis is filled with the line color.       * **Bar Chart** -- Displays each datapoint as a colored vertical bar. Each datapoint has its own bar. The base of the bar is the X-axis, and its top is the datapoint value. Datapoints are closely packed around the X-axis time         value associated with the datapoint. This is also known as a **Column** chart.         <br>         You can optionally display datapoints in a **stacked** bar chart. In this type of chart, different datapoint         values are vertical bars stacked above their associated time. Colors differentiat different datapoints.
	Type string `json:"type"`
	// Specifies the type of unit to use when displaying or labeling values     - Metric\\: Values represent decimal multiples. For example, if `unitPrefix` is Metric, 1K represents 1 kilobyte or 1000       bytes.     - Binary\\: Values represent binary multiples. For example, if `unitPrefix` is Binary, 1K represents 1 kibibyte or 1024       bytes.<br> **Note** This option is available for all values of `option.type` except `Text`
	UnitPrefix string `json:"unitPrefix,omitempty"`
}
