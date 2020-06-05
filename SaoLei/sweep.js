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
        for (let c = 0; c < count; c++) {
            let landmine = {
                x: Math.floor(Math.random() * col),
                y: Math.floor(Math.random() * row)
            };
            if (landmine in landmines) {
                c -= 1;
            }
            else {
                landmines.push(landmine)
            }
        }

        for (let i = 0; i < landmines.length; i++) {
            // 埋雷并划定危险区域
            board[landmines[i].y][landmines[i].x].isLandmine = true;
            board[landmines[i].y][landmines[i].x].isDangerousZone = true;
            this.neighborsHandle(landmines[i].x, landmines[i].y, row, col, board,(neighbor) => {
                neighbor.isDangerousZone = true;
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
    this.isDangerousZone = false;
    this.size = 20;
    this.img = Util.loadImg("images/grid.png");
    this.img.style.height = this.size + 'px';
    this.img.style.width = this.size + 'px';
    this.img.style.left = this.x * this.size + 'px';
    this.img.style.top = this.y * this.size + 'px';
    Board.boardElement.appendChild(this.img)

    this.img.onclick = () => {

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