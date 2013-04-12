function ArticleListCntl($scope, $routeParams, $location, Article) {

	// Get articles
	$scope.articles = [];
	Article.get({}, function(r) {
		$scope.articles = _.map(r.D, function(d) {
			return new Article(d);
		});
	});

	if ($routeParams.articleId) {
		$scope.selected = Article.get({articleId: $routeParams.articleId});
	} else {
		$scope.selected = null;
	}

	$scope.active = function(article) {
		if (article && article == $scope.selected) {
			return "active";
		}

		return "";
	};

	$scope.openArticle = function(article) {
		$scope.selected = article;
	};

	$scope.newArticle = function() {

		var article = new Article({
			title: "New article",
			content: ""			
		});

		article.$save(function(d){
			article._id = d.D._id;
			$scope.articles.push(article)
		});
	};

	$scope.deleteArticle = function() {
		if ($scope.selected) {
			var selected = $scope.selected;
			$scope.selected = null;
			Article.delete({articleId: selected._id}, function(){
				$scope.articles = _.filter($scope.articles, function(article) {
					return selected._id != article._id;
				});
			});
		}
	};
}
