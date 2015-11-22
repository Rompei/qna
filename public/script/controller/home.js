angular.module('qna')
.controller('HomeController', function($scope, $mdDialog, QuestionService){
  var page = 1;
  var maxResults = 40;
  $scope.questions = QuestionService.query({page: page, maxResults: maxResults});

  $scope.prev = function(){
    page--;
    $scope.questions = QuestionService.query({page: page, maxResults: maxResults});
  };

  $scope.next = function(){
    page++;
    $scope.questions = QuestionService.query({page: page, maxResults: maxResults});
  };
  
  $scope.post = function(ev){
    $mdDialog.show({
      controller: PostController,
      templateUrl: 'view/post.html',
      parent: angular.element(document.body),
      parentEvent: ev,
      clickOutsideToClose: true
    })
    .then(function(question){
      QuestionService.save(question, function(){
        $scope.questions = QuestionService.query({page: page, maxResults: maxResults});
      });
    });
  };
  
});

function PostController($scope, $mdDialog){
  $scope.question = {
    title: "",
    selections: [{
      content:""
    }]
  };

  $scope.addSelection = function(){
    $scope.question.selections.push({
      content:""
    })
  };

  $scope.cancel = function(){
    $mdDialog.cancel()
  };

  $scope.post = function(question){
   $mdDialog.hide(question); 
  }
}
