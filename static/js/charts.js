google.charts.load('current', { 'packages': ['line', 'corechart'] });
google.charts.setOnLoadCallback(drawChart);

var dataJson;
var tickVal;
var setTicks = function (month, year) {
    var smonth;
    if (month >= 3) {
        smonth = month - 3;
    } else {
        smonth = 11 - month;
    }
    var m1 = smonth++;
    var m2;
    if ((smonth) < 11) {
        m2 = smonth++;
    } else {
        m2 = smonth = 11;
        smonth++;
    }

    var m3;
    if ((smonth) < 11) {
        m3 = smonth++;
    } else {
        m3 = smonth = 11;
        smonth++;
    }

    var m4;
    if ((smonth) < 11) {
        m4 = smonth++;
    } else {
        m4 = smonth = 11;
        smonth++;
    }
    tickVal = [new Date(year, m1), new Date(year, m2), new Date(year, m3), new Date(year, m4)];    
}

var setChartData = function (val) {
    dataJson = val
}

function drawChart() {

    var button = document.getElementById('change-chart');
    var chartDiv = document.getElementById('chart_div');


    var data = new google.visualization.DataTable(dataJson);
   

    var materialOptions = {
        chart: {
            title: 'Average Latency of API Gateway Per URL for Last Three Months'
        },
        width: 900,
        height: 500,
        series: {
            // Gives each series an axis name that matches the Y-axis below.
            0: { axis: 'Latency' },
            1: { axis: 'Call Count' }
        },
        axes: {
            // Adds labels to each axis; they don't have to match the axis names.
            y: {
                Temps: { label: 'Latency (Milliseconds)' },
                Daylight: { label: 'Call Count' }
            }
        }
    };

    var classicOptions = {
        title: 'Average Latency of API Gateway Per URL',
        width: 900,
        height: 500,
        // Gives each series an axis that matches the vAxes number below.
        series: {
            0: { targetAxisIndex: 0 },
            1: { targetAxisIndex: 1 }
        },
        vAxes: {
            // Adds titles to each axis.
            0: { title: 'Latency (Milliseconds)' },
            1: { title: 'Call Count' }
        },
        hAxis: {
            ticks: tickVal
        },
        vAxis: {
            viewWindow: {
                max: 120
            }
        }
    };

    function drawClassicChart() {
        var classicChart = new google.visualization.LineChart(chartDiv);
        classicChart.draw(data, classicOptions);
    }

    function drawMaterialChart() {
        var materialChart = new google.charts.Line(chartDiv);
        materialChart.draw(data, materialOptions);
    }

    drawClassicChart();

}