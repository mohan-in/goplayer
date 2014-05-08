var app=angular.module('mediaplayer',[]);

function MediaPlayerCtrl($scope) {
	$scope.files = [
		{"name":"song1", "size":"5 MB", "duration":"05:00:00", "src":"path/file", "type":"audio/mpeg"},
		{"name":"song2", "size":"6 MB", "duration":"05:00:00", "src":"path/file", "type":"audio/ogg"},
		{"name":"song3", "size":"7 MB", "duration":"05:00:00", "src":"path/file", "selected":"info"},
		{"name":"song4", "size":"8 MB", "duration":"05:00:00", "src":"path/file"},
		{"name":"song5", "size":"9 MB", "duration":"05:00:00", "src":"path/file"},
		{"name":"song1", "size":"5 MB", "duration":"05:00:00", "src":"path/file"},
		{"name":"song2", "size":"6 MB", "duration":"05:00:00", "src":"path/file"},
		{"name":"song3", "size":"7 MB", "duration":"05:00:00", "src":"path/file"},
		{"name":"song4", "size":"8 MB", "duration":"05:00:00", "src":"path/file"},
		{"name":"song5", "size":"9 MB", "duration":"05:00:00", "src":"path/file"},
		{"name":"song1", "size":"5 MB", "duration":"05:00:00", "src":"path/file"},
		{"name":"song2", "size":"6 MB", "duration":"05:00:00", "src":"path/file"},
		{"name":"song3", "size":"7 MB", "duration":"05:00:00", "src":"path/file"},
		{"name":"song4", "size":"8 MB", "duration":"05:00:00", "src":"path/file"},
		{"name":"song5", "size":"9 MB", "duration":"05:00:00", "src":"path/file"},
	]

	$scope.srcfile = "";

	$scope.play = function(file) {
		console.log("file selected is " + file.src);
		$scope.srcfile = file.src;
		for (var i = $scope.files.length - 1; i >= 0; i--) {
			$scope.files[i].selected = "";
		};
		file.selected = "bg-primary";
	};
};