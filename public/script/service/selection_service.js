angular.module('qna')
.factory('SelectionService', function($resource){
  return $resource('/selections/:id', null, {
    'update': {method:'PUT'}
  });
});
