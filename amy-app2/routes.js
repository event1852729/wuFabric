//SPDX-License-Identifier: Apache-2.0

var cardpayment = require('./controller.js');

module.exports = function(app){

  //app.METHOD(PATH, HANDLER) METHOD 是 HTTP 要求方法。PATH 是伺服器上的路徑。HANDLER 是當路由相符時要執行的函數
  app.get('/get_cardpayment/:id', function(req, res){
    cardpayment.get_cardpayment(req, res);
  });
  app.get('/add_cardpayment/:cardpayment', function(req, res){
    cardpayment.add_cardpayment(req, res);
  });
  app.get('/get_all_cardpayment', function(req, res){
    cardpayment.get_all_cardpayment(req, res);
  });
  app.get('/change_value/:newvalue', function(req, res){
    cardpayment.change_value(req, res);
  });
  app.get('/delete_cardpayment/:id', function(req, res){
    cardpayment.delete_cardpayment(req, res);
  });
  app.get('/addfivedata', function(req, res){
    cardpayment.addfivedata(req, res);
  });
}
