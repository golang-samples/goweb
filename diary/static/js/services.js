angular.module('diaryServices', ['ngResource']).
	factory('Article', function($resource) {
		return $resource('/api/articles/:articleId', {}, {});
	});
