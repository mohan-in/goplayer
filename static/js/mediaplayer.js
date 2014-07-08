var app=angular.module('mediaplayer',['ui.bootstrap']);

app.c

app.directive("filelist", function() {
	return {
		scope: true,
		restrict: 'E',
      	replace: 'true',
    	templateUrl: './filelist.html',
    	controller: function ($scope, $http, $rootScope) {
			$scope.srcfile = "";
			$scope.breadcrumb = [{Name: "/...", AbsPath: ""}];

		    HttpGet("/");
		    ishover=false;
			$scope.play = function(file) {
				if(file.IsDir === true) {
					$scope.breadcrumb.push({Name: file.Name, AbsPath: file.AbsPath});
					HttpGet("/?path=" + file.AbsPath);
				} else {
					$rootScope.audioSrc = "/media/?file=" + file.AbsPath;
					for (var i = $scope.files.length - 1; i >= 0; i--) {
						$scope.files[i].isPlaying = false;
					};
					file.isPlaying = true;
				}	
			};

			$scope.gotoCrumb = function(i, p) {
				$scope.breadcrumb = $scope.breadcrumb.slice(0, i + 1);
				HttpGet("/?path=" + p);
			};

			function HttpGet(path) {
				$http.get(path).success(function(data) {	
					$scope.files = data;
				});
			}
		}	
    };
});

