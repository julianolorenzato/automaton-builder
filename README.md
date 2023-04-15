# Automaton Builder


<!-- ## ABDLang

ABDLang (Automaton Builder Description Language) is a language to describe the automatons very inspired on DOT Language. -->

![](./example.png)


```
{
  "name": "Deterministic Finite Automaton",
  "alphabet": ["0", "1"],
  "states": [
    {
      "name": "q0",
      "initial": true,
      "final": true
    },
    {
      "name": "q1",
      "initial": false,
      "final": true
    },
    {
      "name": "q2",
      "initial": false,
      "final": true
    }
  ],
  "transitions": [
    {
      "from": "q0",
      "to": "q0",
      "symbol": "0"
    },
    {
      "from": "q0",
      "to": "q1",
      "symbol": "1"
    },
    {
      "from": "q1",
      "to": "q0",
      "symbol": "0"
    },
    {
      "from": "q1",
      "to": "q2",
      "symbol": "1"
    },
    {
      "from": "q2",
      "to": "q0",
      "symbol": "0"
    }
  ]
}

```
[example.json](./example.json)
