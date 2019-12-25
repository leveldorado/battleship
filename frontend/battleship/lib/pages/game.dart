


import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class Game extends StatefulWidget {
   Field field;
   Game();
  _GameState createState() => _GameState();
}

class _GameState extends State<Game> {

  Future<Field> _createGame() async {
    return Field();
  }

  @override
  void initState() {
    _createGame().then((Field field){
      widget.field = field;
    });
    super.initState();
  }


  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.white,
      child: Center(),
    );
  }
}


class Field {
  List<List<Cell>> cells;
}


class Cell {
  CellState state;
}


enum CellState {
  hit,
  miss,
  ship,
  empty,
}