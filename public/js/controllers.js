var jfmControllers = angular.module('jfmControllers', []);

jfmControllers.controller('jfmUpload', ['$scope', '$http',
	function($scope, $http) {
		$scope.uploaded = false;
		$scope.results = {};

		$scope.processUpload = function(files) {
			console.log("Processing...");
			var fd = new FormData();
			fd.append("csv", files[0]);

			$http.post('upload', fd, {
				headers : { 'Content-Type': undefined },
				transformRequest: angular.identity
			})
			.success(function(response) {
					console.log(response);

					$scope.uploaded = true;
					$scope.results = response;
					// Extra result for demonstration
					// $scope.results[1] = {
					// 	"Name": "Susan",
					// 	"Total": 110.25
					// }
				})
			.error(function() {
				console.log("Something failed");
			});
		}
	}]);