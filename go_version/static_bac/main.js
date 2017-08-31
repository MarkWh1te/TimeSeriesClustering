/** * Created by Ky on 2017/8/13. */
var c = function (index, num, data, data_2, data_3) {
    var id = "chart_" + index
    var arr = []
    var len = data.length
    while (len) {

        arr.push(len)
        len--
    }
    var series = []
    for (var i in data) {
        series.push(
            {
                "name": data_2[data[i]],
                "data": data_3[data_2[data[i]]]
            }
        )
    }
    console.log(series)
    var chart = new Highcharts.Chart(id, {// 图表初始化函数，其中 container 为图表的容器 div
        chart: {
            type: 'line'                           //指定图表的类型，默认是折线图（line）
        },
        title: {
            text: '分类:' + index//指定图表标题
        },
        xAxis: {
            categories: len   //日期
        },
        yAxis: {
            title: {
                text: 'something'                 //指定y轴的标题
            }
        },
        series: series
    });
}


jQuery(document).ready(function () {
     /* attach a submit handler to the form */
    $("#searchForm").submit(function(event) {
        /* stop form from submitting normally */
        event.preventDefault();
        /* get some values from elements on the page: */
        var $form = $(this),
            types = $form.find('input[name="types"]').val()
            iters = $form.find('input[name="iters"]').val()
        console.log(types,iters)

        /* Send the data using post */
     $.ajax({
        url: "/cluster",
        type: "post",
        async: false,
        //data: {"days": 5, "types": 3},
        data: {"days": types, "types": iters},
        success: function (data) {
            console.log(data)
            clusters = data["Cluster"]
            // alert(clusters)
            for (var i in clusters) {
                console.log(i)
        /* Put the results in a div */
                $("#chart_area").append('<li id="chart_' + i + '"></li>')
                c(i, 10, clusters[i], data["Sort_keys"], data["Source"])
            }
        }
    })           

    });
    $.ajax({
        url: "/cluster",
        type: "post",
        async: false,
        //data: {"days": 5, "types": 3},
        data: {"days": 5, "types": 6},
        success: function (data) {
            console.log(data)
            clusters = data["Cluster"]
            // alert(clusters)
            for (var i in clusters) {
                console.log(i)
                $("#chart_area").append('<li id="chart_' + i + '"></li>')
                c(i, 10, clusters[i], data["Sort_keys"], data["Source"])
            }
        }
    })
});
