angular.module('diary', ['diaryServices']).
	config(['$routeProvider', function($routeProvider) {
		$routeProvider.
			when('/articles/:articleId', 
				{
					templateUrl: 'index.html',
					controller: ArticleListCntl
				}).
			otherwise({redirectTo: '/articles/'});
	}]);
