var app=angular.module('mediaplayer',[]);

function MediaPlayerCtrl($scope, $http) {

	$http.get("/").
    	success(function(data) {
    		$scope.files = data;
    	});

	// $scope.play = function(file) {
	// 	console.log("file selected is " + file.src);
	// 	$scope.srcfile = file.src;
	// 	for (var i = $scope.files.length - 1; i >= 0; i--) {
	// 		$scope.files[i].selected = "";
	// 	};
	// 	file.selected = "bg-primary";
	// };
};