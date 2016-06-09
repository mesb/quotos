/*
  Angular Js services

*/

var quotosApp = angular.module("quotosApp", []).config(["$interpolateProvider", function($interpolateProvider){
                $interpolateProvider.startSymbol("{$");
                $interpolateProvider.endSymbol("$}");
}]);



quotosApp.factory('QuotesService', function($http){
    quotesFactory= {};

    // link to the quotes api url
    var quotes_url= '/list';
    var post_url= '/snippet';

    // service that returns quotes from api
    quotesFactory.list= function(){
        return $http.get(quotes_url);
    };

    quotesFactory.poster = function(snippet) {
        return $http.post(post_url, snippet);
    };

    return quotesFactory;
});

quotosApp.controller('QuotesController', function($scope, $timeout, $interval, QuotesService) {


    $scope.QUOTE = {
        body: '',
        submitter: '',
        author: ''
    }

    // post quote to server
    $scope.postSnippet= function() {


        var p= QuotesService.poster($scope.QUOTE);
        p
            .then(function(response){
                console.log(response.data);
                $scope.quotes.unshift($scope.QUOTE);
                $scope.msgPane= response.data.Body.success;

                $timeout(function() {
                    $scope.msgPane= "";
                }, 5000);

                $scope.clearQUOTE();
            }, function(err){
                console.log(err.data.Body);
                var dat= err.data.Body;
                var txt= " ";

                for (var prop in dat) {
                    txt += dat[prop] + "; ";
                }
                $scope.msgPane= txt;
            });

    };


    $scope.clearQUOTE = function() {
        $scope.QUOTE = {
        body: '',
        submitter: '',
        author: ''
        };
    };


    $scope.clear= function(){
        $scope.quotes= [];
    };

/*
    $interval( function(){ $scope.quotes.push({"body":"bwahahahahather you go
    /*up the ladder or down the ladder, your position is shaky",
    /*"submitter":"mesb", "author": "Lao Tzu"});}, 3000000);

    */

    var promise= QuotesService.list();

    promise
        .then(function(response){

            $scope.quotes= response.data;
        });

});
