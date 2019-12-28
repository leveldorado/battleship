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
      super.setState((){
        widget.field = field;
      });
    });
    super.initState();
  }

  _fire(int x, int y) {
    print("pressed $x and $y");
  }

  Column _buildColumn() {
    if (widget.field == null) {
      return Column();
    }
    List<Row> rows = new List();
    for (var i = 0; i < widget.field.height; i++) {
      List<IconButton> buttons = new List();
      for (var j = 0; j < widget.field.width; j++) {
        buttons.add(IconButton(
          icon: Icon(Icons.check_box_outline_blank),
          onPressed: () => _fire(j, i),
        ));
      }
      rows.add(Row(children: buttons));
    }
    for (var ship in widget.field.ships) {
       for (var coordinate in ship.coordinates) {
         var x = rows.length;
         rows[coordinate.y].children[coordinate.x] = IconButton(
           icon: Icon(Icons.toys),
           onPressed: () => _fire(coordinate.x, coordinate.y),
         );
       }
    }
    return Column(children: rows);
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
