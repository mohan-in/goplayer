angular.module('mediaplayer',['ui.bootstrap'])

.directive("filelist", function() {
	return {
		scope: true,
		restrict: 'E',
      	replace: 'true',
    	templateUrl: 'static/filelist.html',
    	controller: function ($scope, $http, $rootScope) {
			$scope.breadcrumb = [{Name: "/...", AbsPath: ""}];

		    HttpGet("/dir");
		    ishover=false;
			$scope.play = function(file) {
				if(file.IsDir === true) {
					$scope.breadcrumb.push({Name: file.Name, AbsPath: file.AbsPath});
					HttpGet("/dir?path=" + file.AbsPath);
				} else if(file.IsAudio) {
					$rootScope.audioSrc = "/media/?file=" + file.AbsPath;
					$rootScope.audiosrcName = file.Name;
					for (var i = $scope.files.length - 1; i >= 0; i--) {
						$scope.files[i].isPlaying = false;
					};
					file.isPlaying = true;
				}	
			};

			$scope.gotoCrumb = function(i, p) {
				$scope.breadcrumb = $scope.breadcrumb.slice(0, i + 1);
				HttpGet("/dir?path=" + p);
			};

			function HttpGet(path) {
				$http.get(path).success(function(data) {	
					$scope.files = data;
				});
			}
		}	
    };
});

