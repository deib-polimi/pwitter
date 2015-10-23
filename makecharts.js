$(document).ready( function(){



var GREEN = function(alpha){ return "rgba(15,157,88," + alpha + ")"; };
var RED = function(alpha){ return "rgba(219,68,55," + alpha + ")"; };

var makechart = function(data, id){
    var cpu_ctx = $("#" + id + "-cpu").get(0).getContext("2d");
    var mem_ctx = $("#" + id + "-mem").get(0).getContext("2d");
    var cpu_ch = new Chart(cpu_ctx);
    var mem_ch = new Chart(mem_ctx);
    var spacing = 0.1;

    var cpu_points = {
        labels: [],
        datasets: [
            {
                label: 'cpu',
                fillColor: RED(0.2),
                strokeColor: RED(1),
                pointColor: RED(1),
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: RED(1),
                data: [],
            },
        ]
    };

    var mem_points = {
        labels: [],
        datasets: [
            {
                label: 'memory',
                fillColor: GREEN(0.2),
                strokeColor: GREEN(1),
                pointColor: GREEN(1),
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: GREEN(1),
                data: [],
            },
        ]
    };

    for(var i = 0; i < data.length; i++){
        var cpu = data[i]['tot']['cpu_percent'];
        var mem = data[i]['tot']['memory_percent'];
        cpu_points['labels'].push((spacing * i).toFixed(1));
        mem_points['labels'].push((spacing * i).toFixed(1));
        cpu_points['datasets'][0]['data'].push(cpu);
        mem_points['datasets'][0]['data'].push(mem);
    }

    cpu_ch.Line(cpu_points);
    mem_ch.Line(mem_points);
};


$.getJSON(
    "500get.json",
    function(data){
        makechart(data, 'g');
    }
);

$.getJSON(
    "500post.json",
    function(data){
        makechart(data, 'p');
    }
);

$.getJSON(
    "250get250post.json",
    function(data){
        makechart(data, 'gp');
    }
);


});
