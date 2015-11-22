angular.module('qna')
.controller('QuestionController', function($scope, $stateParams, QuestionService, SelectionService, CommentService){
  $scope.question = QuestionService.get({id :$stateParams.id});
  
  $scope.vote = function(selection){
    selection.count++;
    SelectionService.update({ id: selection.id }, selection);
  };

  $scope.post = function(){
    $scope.newComment.questionId = $scope.question.id;
    CommentService.save($scope.newComment, function(){
      $scope.question = QuestionService.get({id :$stateParams.id});
    });
  };
});
