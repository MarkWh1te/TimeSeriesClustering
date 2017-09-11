/**
 * Created by Ky on 2017/8/13.
 */

// 日期格式化方法
Date.prototype.format = function (format) {
    var date = {
        "M+": this.getMonth() + 1,
        "d+": this.getDate(),
        "h+": this.getHours(),
        "m+": this.getMinutes(),
        "s+": this.getSeconds(),
        "q+": Math.floor((this.getMonth() + 3) / 3),
        "S+": this.getMilliseconds()
    };
    if (/(y+)/i.test(format)) {
        format = format.replace(RegExp.$1, (this.getFullYear() + '').substr(4 - RegExp.$1.length));
    }
    for (var k in date) {
        if (new RegExp("(" + k + ")").test(format)) {
            format = format.replace(RegExp.$1, RegExp.$1.length == 1
                ? date[k] : ("00" + date[k]).substr(("" + date[k]).length));
        }
    }
    return format;
}
var c = function (char_id,index, num, data, data_2, data_3) {
    var id = char_id + index
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
    var request_date = function (args) {
        $.ajax({
            // url: "http://127.0.0.1:5000/cluster", 
            url: "http://127.0.0.1:8080/cluster",
            type: "post",
            async: false,
            data: args,
            success: function (data) {
                console.log(data)
                clusters = data["Cluster"]
                // alert(clusters)
                $("#chart_area").html("")
                for (var i in clusters) {
                    console.log(i)
                    $("#chart_area").append('<li id="chart_' + i + '"></li>')
                    // $("#chart_area2").append('<li id="chart2_' + i + '"></li>')
                    c("chart_",i, 10, clusters[i], data["Sort_keys"], data["Source"])
                    // c("chart2_",i, 10, clusters[i], data["Sort_keys"], data["Origin"])
                }
            }
        })
    }
    $(".ui-select").selectWidget({
        change: function (changes) {
            return changes;
        },
        effect: "slide",
        keyControl: true,
        speed: 200,
        scrollHeight: 250
    });
    $('.ui-choose').ui_choose();
    var uc_04 = $('#uc_04').data('ui-choose');
    var stocks = "3"
    uc_04.click = function (index, item) {
        stocks = index
        stocks = stocks.join(",")
    };
    $.datepicker.regional['zh-CN'] = {
        clearText: '清除',
        clearStatus: '清除已选日期',
        closeText: '关闭',
        closeStatus: '不改变当前选择',
        prevText: '< 上月',
        prevStatus: '显示上月',
        prevBigText: '<<',
        prevBigStatus: '显示上一年',
        nextText: '下月>',
        nextStatus: '显示下月',
        nextBigText: '>>',
        nextBigStatus: '显示下一年',
        currentText: '今天',
        currentStatus: '显示本月',
        monthNames: ['一月', '二月', '三月', '四月', '五月', '六月', '七月', '八月', '九月', '十月', '十一月', '十二月'],
        monthNamesShort: ['一月', '二月', '三月', '四月', '五月', '六月', '七月', '八月', '九月', '十月', '十一月', '十二月'],
        monthStatus: '选择月份',
        yearStatus: '选择年份',
        weekHeader: '周',
        weekStatus: '年内周次',
        dayNames: ['星期日', '星期一', '星期二', '星期三', '星期四', '星期五', '星期六'],
        dayNamesShort: ['周日', '周一', '周二', '周三', '周四', '周五', '周六'],
        dayNamesMin: ['日', '一', '二', '三', '四', '五', '六'],
        dayStatus: '设置 DD 为一周起始',
        dateStatus: '选择 m月 d日, DD',
        dateFormat: 'yy-mm-dd',
        firstDay: 1,
        initStatus: '请选择日期',
        isRTL: false
    };
    $.datepicker.setDefaults($.datepicker.regional['zh-CN']);
    $(function () {
        $("#start_date").datepicker(
            {
                changeMonth: true,
                changeYear: true,
                yearRange: '2011:2017',
                maxDate: -1,
                dateFormat: 'yymmdd'
            }
        );
        $("#start_date").datepicker("setDate", "20170715")
        $("#end_date").datepicker(
            {
                changeMonth: true,
                changeYear: true,
                yearRange: '2011:2017',
                maxDate: 0,
                dateFormat: 'yymmdd'
            }
        );
        $("#start").on("click", function () {
            // alert($("#types").html())
            args = {

                "start_date": Number($("#start_date").datepicker('getDate').format("yyyyMMdd")),
                "end_date": Number($("#end_date").datepicker('getDate').format("yyyyMMdd")),
                "stock": stocks,
                "types": $("#types").val(),
                "method": $("#method").val(),
                "algorithms": $("#algorithms").val()
            }
            console.log(args)
            request_date(args)
        })
        $("#end_date").datepicker("setDate", "20170731")
        request_date({
            "start_date": 20170715,
            "end_date": 20170731,
            "stock": "3",
            // "stock": "0,6,3",
            "method": 0,
            "types": 5,
            "algorithms":1
        })
    });

});
