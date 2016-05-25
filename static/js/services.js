/*
  Angular Js services

*/

var app = angular.module("quotos", []).config(["$interpolateProvider", function($interpolateProvider){
                $interpolateProvider.startSymbol("{$");
                $interpolateProvider.endSymbol("$}");
           }]);

app.controller('QuotesCtrl', function($scope, $http) {
    $http.get("/quotes/quote.json")
    .then(function(response) {
        //First function handles success
        $scope.quotes = response.data.quotes;

    }, function(response) {
        //Second function handles error
        $scope.content = "Something went wrong";
    });
});
