angular.module('qna', 
    [
      'ngMaterial',
      'ngResource',
      'ui.router'
    ])
.config(function($stateProvider, $urlRouterProvider){
  $urlRouterProvider.otherwise('/');
  
  $stateProvider
    .state('home', {
      url: '/',
      templateUrl: 'view/home.html',
      controller: 'HomeController'
    })
    .state('question', {
      url:'/question/:id',
      templateUrl: 'view/question.html',
      controller: 'QuestionController'
    });
});
