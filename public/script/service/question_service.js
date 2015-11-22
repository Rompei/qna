angular.module('qna')
.factory('QuestionService', function($resource){
  return $resource('/questions/:id', null, {
    'update': {method:'PUT'}
  });
});
  
