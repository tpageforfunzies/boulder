$(document).ready(function() {
        $.get('/v1/logs/gin.log', function(data) {

		   $("#logs").text(data);
		   updateTime()

		}, 'text');
});

$(document).ready(function() {
    $("#logButton").click(function(){
        $.get('/v1/logs/gin.log', function(data) {

		   $("#logs").text(data);
		   updateTime()

		}, 'text');
    }); 
});

function updateTime() {
	var d = new Date()
	$("#lastUpdated").text("Last Updated: " + d.toUTCString())
}