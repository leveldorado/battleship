



import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class StartPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        // Here we take the value from the MyHomePage object that was created by
        // the App.build method, and use it to set our appbar title.
        title: Text("Start battleship game"),
      ),
      body: Center(
        child: RaisedButton(
          onPressed: () {},
          child: const Text(
              'Start',
              style: TextStyle(fontSize: 20)
          ),
        ),
      ),
    );
  }
}