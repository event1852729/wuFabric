// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();
	
	$scope.queryAllCardPayment = function(){

		appFactory.queryAllCardPayment(function(data){
			var array = [];
			for (var i = 0; i < data.length; i++){
				parseInt(data[i].Key);
				data[i].Record.Key = parseInt(data[i].Key);
				array.push(data[i].Record);
				console.log(data[i].Record);
			}
			array.sort(function(a, b) {
			    return parseFloat(a.Key) - parseFloat(b.Key);
			});
			$scope.all_cardpayment = array;
		});
	}

	$scope.queryCardPayment = function(){

		var id = $scope.cardpayment_id;

		appFactory.queryCardPayment(id, function(data){
			$scope.query_cardpayment = data;

			if ($scope.query_cardpayment == "Could not locate tuna"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	$scope.recordCardPayment = function(){

		appFactory.recordCardPayment($scope.cardpayment, function(data){
			$scope.create_tuna = data;
			$("#success_create").show();
		});
	}

	$scope.changeValue = function(){

		appFactory.changeValue($scope.value, function(data){
			$scope.change_value = data;
			if ($scope.change_value == "Error: no tuna catch found"){
				$("#error_holder").show();
				$("#success_holder").hide();
			} else{
				$("#success_holder").show();
				$("#error_holder").hide();
			}
		});
	}

});

// Angular Factory
app.factory('appFactory', function($http){
	
	var factory = {};

    factory.queryAllCardPayment= function(callback){

    	$http.get('/get_all_cardpayment/').success(function(output){
			callback(output)
		});
	}

	factory.queryCardPayment = function(id, callback){
    	$http.get('/get_cardpayment/'+id).success(function(output){
			callback(output)
		});
	}

	factory.recordCardPayment = function(data, callback){


		var cardpayment = data.id + "-" + data.acquiringbank + "-" + data.acquiringdate + "-" + data.size + "-" + data.ipfsHash;

    	$http.get('/add_cardpayment/'+cardpayment).success(function(output){
			callback(output)
		});
	}

	factory.changeValue = function(data, callback){

		var newvalue = data.id + "-" + data.newsize + "-" + data.newipfshash;

    	$http.get('/change_value/'+newvalue).success(function(output){
			callback(output)
		});
	}

	return factory;
});


