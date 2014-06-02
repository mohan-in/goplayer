var app=angular.module('mediaplayer',[]);

function MediaPlayerCtrl($scope, $http) {

	$http.get("/").success(function(data) {
    		$scope.files = data;
    });

	$scope.srcfile = "";

	$scope.play = function(file) {
		console.log("file selected is " + file.AbsPath);
		if(file.IsDir === true) {
			$http.get("/?path="+file.AbsPath).success(function(data) {
				$scope.files = data;
			});
		} else {
			$scope.srcfile = "/media/?file=" + file.AbsPath;
			file.isPlaying = true;
		}
		
	};
};