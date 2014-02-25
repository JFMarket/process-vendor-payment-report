var jfm = angular.module('jfm', [
	'ngRoute',
	'jfmControllers']);

jfm.config(['$routeProvider',
	function($routeProvider) {
		$routeProvider.
			when('/', {
				templateUrl: 'partials/index.html',
				controller: 'jfmUpload'
			}).
			otherwise({
				redirectTo: '/'
			});
	}]);