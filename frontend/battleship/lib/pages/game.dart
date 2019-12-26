import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:global_configuration/global_configuration.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

const _createGamePath = "/api/v1/game";

class Game extends StatefulWidget {
  Field field;
  Game();
  _GameState createState() => _GameState();
}

class _GameState extends State<Game> {
  final host = GlobalConfiguration().getString("host");
  Future<Field> _createGame() async {
    final response = await http.post('http://{$host}{$_createGamePath}');
    return Field.fromJson(json.decode(response.body));
  }

  @override
  void initState() {
    _createGame().then((Field field) {
      widget.field = field;
    });
    super.initState();
  }

  Column _buildColumn() {
    List<Row> rows;
    for (var row in widget.field.cells) {
        List<IconButton> buttons;
        for (var cell in row) {
          var icon = Icons.check_box_outline_blank;
          switch (cell.state) {
            case CellStateMiss:
              icon = Icons.adjust;
              break;
            case CellStateHit:
              icon = Icons.clear;
              break;
            case CellStateShip:
              icon = Icons.toys;

          }
          buttons.add(IconButton(
            icon: Icon(icon),
            onPressed: (){
              print("pressed $cell.x and $cell.y");
            },
          ));
        }
        rows.add(Row(children: buttons));
    }
    return Column(children: rows);
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.white,
      child: Center(
        child: _buildColumn(),
      ),
    );
  }
}

class Field {
  List<List<Cell>> cells;

  Field({this.cells});

  factory Field.fromJson(Map<String, dynamic> json) {
    var cells = json["cells"];
    return Field(
      cells: cells,
    );
  }
}

const CellStateHit = "hit";
const CellStateMiss = "miss";
const CellStateShip = "ship";

class Cell {
  String state;
  int x;
  int y;
}
