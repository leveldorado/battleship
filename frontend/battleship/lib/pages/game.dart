import 'dart:io';

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
    final response = await http.post('http://$host$_createGamePath');
    print(response.body);
    return Field.fromJson(json.decode(response.body));
  }

  @override
  void initState() {
    _createGame().then((Field field) {
      super.setState(() {
        widget.field = field;
      });
    });
    super.initState();
  }

  _fire(int x, int y) {
    print("pressed $x and $y");
  }

  Widget _buildColumn() {
    if (widget.field == null) {
      return Column();
    }
    List<Row> rows = new List();
    var screenSize = MediaQuery.of(context).size;
    var width = screenSize.width;
    if (MediaQuery.of(context).orientation == Orientation.landscape) {
        width = screenSize.height * 0.8;
    }
    var cellWidth = width / widget.field.width;
    var cellHeight = width / widget.field.height;
    for (var i = 0; i < widget.field.height; i++) {
      List<Widget> buttons = new List();
      for (var j = 0; j < widget.field.width; j++) {
        buttons.add(Ink(
          width: cellWidth,
          height: cellHeight,
          decoration: BoxDecoration(border: Border.all()),
          child: InkWell(
            onTap: () => _fire(j, i),
          ),
        ));
      }
      rows.add(Row(children: buttons));
    }
    for (var ship in widget.field.ships) {
      for (var coordinate in ship.coordinates) {
        rows[coordinate.y].children[coordinate.x] = Ink(
          width: cellWidth,
          height: cellHeight,
          decoration: BoxDecoration(border: Border.all(),
              image: DecorationImage(
                  image: AssetImage("assets/images/ship.jpeg"),
                fit: BoxFit.fill,
              )
          ),
          child: InkWell(
            onTap: () => _fire(coordinate.x, coordinate.y),
          ),
        );
      }
    }
    var column = Column(children: rows);
    if (MediaQuery.of(context).orientation == Orientation.landscape) {
      return Container(padding: EdgeInsets.only(left: width * 0.2),child: column);
    }
    return column;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Game"),
      ),
      body: Center(
        child: _buildColumn(),
      ),
    );
  }
}

class Field {
  List<Ship> ships;
  int width;
  int height;

  Field({this.ships, this.width, this.height});

  factory Field.fromJson(Map<String, dynamic> json) {
    var list = json["Ships"] as List;
    return Field(
      ships: list.map((i) => Ship.fromJson(i)).toList(),
      width: json["Width"],
      height: json["Height"],
    );
  }
}

class Ship {
  List<Coordinate> coordinates;
  Ship({this.coordinates});
  factory Ship.fromJson(Map<String, dynamic> json) {
    var list = json['Coordinates'] as List;
    return Ship(coordinates: list.map((i) => Coordinate.fromJson(i)).toList());
  }
}

class Coordinate {
  int x;
  int y;

  Coordinate({this.x, this.y});

  factory Coordinate.fromJson(Map<String, dynamic> json) {
    return Coordinate(x: json["X"], y: json["Y"]);
  }
}
