let Main = {
    row: 20,
    col: 15,
    count: 30,
    start: function () {
        this.board = Board.create(this.row, this.col);
        this.board = Board.bury(this.count, this.row, this.col, this.board)
    }

};
let Board = {
    create: function (row, col) {
        this.boardElement = document.getElementById("board");
        this.boardElement.style.position = "relative";
        this.boardElement.style.height = row * 20 + 'px';
        this.boardElement.style.width = col * 20 + 'px';

        let board = [];
        for (let y = 0; y < row; y++) {
            let _col = [];
            board.push(_col);
            for (let x = 0; x < col; x++) {
                let grid = new Grid(x, y);
                board[y].push(grid);
            }
        }
        return board
    },
    bury: function (count, row, col, board) {
        let landmines = [];
        let col_index = col - 1;
        let row_index = row - 1;
        for (let c = 0; c < count; c++) {
            let landmine = {
                x: Math.floor(Math.random() * col_index),
                y: Math.floor(Math.random() * row_index)
            };
            if (landmine in landmines) {
                c -= 1;
            }
            else {
                landmines.push(landmine)
            }
        }

        for (let i = 0; i < landmines.length; i++) {
            // 埋雷并划定周边范围为危险区域
            board[landmines[i].y][landmines[i].x].isLandmine = true;
            // board[landmines[i].y][landmines[i].x].img.src = "images/landmine.png";
            this.neighborsHandle(landmines[i].x, landmines[i].y, row_index, col_index, board, (neighbor) => {
                neighbor.neighborLandminesNumber += 1;
                // if (!(neighbor.isLandmine)) {
                //     neighbor.img.src = "images/" + neighbor.neighborLandminesNumber + ".png"
                // }
            });
        }

        return board
    },
    neighborsHandle: function (x, y, row, col, board, action) {
        let neighbors = [];
        if (x > 0) {
            neighbors.push(board[y][x - 1]);
        }
        if (y > 0) {
            neighbors.push(board[y - 1][x]);
        }
        if (x < col) {
            neighbors.push(board[y][x + 1]);
        }
        if (y < row) {
            neighbors.push(board[y + 1][x]);
        }
        if (x > 0 && y > 0) {
            neighbors.push(board[y - 1][x - 1]);
        }
        if (x < col && y < row) {
            neighbors.push(board[y + 1][x + 1]);
        }
        if (y > 0 && x < col) {
            neighbors.push(board[y - 1][x + 1]);
        }
        if (x > 0 && y < row) {
            neighbors.push(board[y + 1][x - 1]);
        }
        for (let j = 0; j < neighbors.length; j++) {
            action(neighbors[j])
        }
    }

};

function Grid(x, y) {
    this.x = x;
    this.y = y;
    this.isLandmine = false;
    this.neighborLandminesNumber = 0;
    this.size = 20;
    this.img = Util.loadImg("images/grid.png");
    this.img.style.height = this.size + 'px';
    this.img.style.width = this.size + 'px';
    this.img.style.left = this.x * this.size + 'px';
    this.img.style.top = this.y * this.size + 'px';
    Board.boardElement.appendChild(this.img);
    this.img.oncontextmenu = () => {
        // 屏蔽右键菜单(contextmenu)
        return false
    };
    this.img.onmousedown = (event) => {
        let btnNum = event.button;
        if (btnNum == 0) {
            //点击鼠标左键
            if (this.isLandmine) {
                this.img.src = "images/landmine.png";
                alert("结束");
            }
            else if (this.neighborLandminesNumber > 0) {
                this.img.src = "images/" + this.neighborLandminesNumber + ".png";
            }
            else {
                this.img.src = "images/blank.png";
                function showNeighors(neighbor) {
                    if (neighbor.neighborLandminesNumber > 0) {
                        neighbor.img.src = "images/" + neighbor.neighborLandminesNumber + ".png";
                        return
                    }
                    else if (!(neighbor.isLandmine)){
                        neighbor.img.src = "images/blank.png";
                    }
                    else {
                        return
                    }
                    Board.neighborsHandle(neighbor.x, neighbor.y, Main.row-1, Main.col-1, Main.board, showNeighors)
                };
                Board.neighborsHandle(this.x, this.y, Main.row-1, Main.col-1, Main.board, showNeighors)
            }
        }
        else if (btnNum == 2) {
            //点击鼠标右键
            this.img.src = "images/flag.png";
        }

    }
};

let Util = {
    loadImg: function (src) {
        let e = document.createElement("img");
        e.style.position = "absolute";
        e.src = src;
        return e;
    }
};