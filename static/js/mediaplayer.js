var app=angular.module('mediaplayer',[]);

function MediaPlayerCtrl($scope, $http) {

	$scope.srcfile = "";
	$scope.breadcrumb = [{Name: "/...", AbsPath: ""}];

    HttpGet("/");


	$scope.play = function(file) {
		if(file.IsDir === true) {
			$scope.breadcrumb.push({Name: file.Name, AbsPath: file.AbsPath});
			HttpGet("/?path=" + file.AbsPath);
		} else {
			$scope.srcfile = "/media/?file=" + file.AbsPath;
			file.isPlaying = true;
		}	
	};

	$scope.gotoCrumb = function(i, p) {
		console.log(i);
		$scope.breadcrumb = $scope.breadcrumb.slice(0, i + 1);
		HttpGet("/?path=" + p);
	};

	function HttpGet(path) {
		$http.get(path).success(function(data) {
			$scope.files = data;
		});
	}
};