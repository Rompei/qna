angular.module('qna')
.factory('CommentService', function($resource){
  return $resource('/comments/:id', null, {
    'update': {method:'PUT'}
  });
});
