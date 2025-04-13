import numpy as np
import pygame
import sys
import math
import random

ROWS = 7
COLS = 7

PLAYER_TURN = 0
AI_TURN = 1

PLAYER_PIECE = 0
AI_PIECE = 2

BLUE = (0, 0, 255)
BLACK = (0, 0, 0)
RED = (255, 0, 0)
YELLOW = (255, 255, 0)


def create_board():
    board = np.zeros((ROWS, COLS))
    board = board + 1
    return board

def drop_piece(board, row, col, piece):
    board[row][col] = piece

def is_valid_location(board, col):
    return board[0][col] == 1

def get_next_open_row(board, col):
    for r in range(ROWS-1, -1, -1):
        if board[r][col] == 1:
            return r

def winning_move(board, piece):
    for c in range(COLS-3):
        for r in range(ROWS):
            if board[r][c] == piece and board[r][c+1] == piece and board[r][c+2] == piece and board[r][c+3] == piece:
                return True

    for c in range(COLS):
        for r in range(ROWS-3):
            if board[r][c] == piece and board[r+1][c] == piece and board[r+2][c] == piece and board[r+3][c] == piece:
                return True

    for c in range(COLS-3):
        for r in range(3, ROWS):
            if board[r][c] == piece and board[r-1][c+1] == piece and board[r-2][c+2] == piece and board[r-3][c+3] == piece:
                return True

    for c in range(3,COLS):
        for r in range(3, ROWS):
            if board[r][c] == piece and board[r-1][c-1] == piece and board[r-2][c-2] == piece and board[r-3][c-3] == piece:
                return True

def draw_board(board):
    for c in range(COLS):
        for r in range(ROWS):
            pygame.draw.rect(screen, BLUE, (c * SQUARESIZE, r * SQUARESIZE + SQUARESIZE, SQUARESIZE, SQUARESIZE ))
            if board[r][c] == 1:
                pygame.draw.circle(screen, BLACK, (int(c * SQUARESIZE + SQUARESIZE/2), int(r* SQUARESIZE + SQUARESIZE + SQUARESIZE/2)), circle_radius)
            elif board[r][c] == 0:
                pygame.draw.circle(screen, RED, (int(c * SQUARESIZE + SQUARESIZE/2), int(r* SQUARESIZE + SQUARESIZE + SQUARESIZE/2)), circle_radius)
            else :
                pygame.draw.circle(screen, YELLOW, (int(c * SQUARESIZE + SQUARESIZE/2), int(r* SQUARESIZE + SQUARESIZE + SQUARESIZE/2)), circle_radius)

    pygame.display.update()


def evaluate_window(window, piece):
    opponent_piece = PLAYER_PIECE
    if piece == PLAYER_PIECE:
        opponent_piece = AI_PIECE
    score = 0
    if window.count(piece) == 4:
        score += 100
    elif window.count(piece) == 3 and window.count(1) == 1:
        score += 5
    elif window.count(piece) == 2 and window.count(1) == 2:
        score += 2
    if window.count(opponent_piece) == 3 and window.count(1) == 1:
        score -= 4 

    return score    

def score_position(board, piece):
    score = 0
    center_array = [int(i) for i in list(board[:,COLS//2])]
    center_count = center_array.count(piece)
    score += center_count * 6
    for r in range(ROWS):
        row_array = [int(i) for i in list(board[r,:])]
        for c in range(COLS - 3):
            window = row_array[c:c + 4]
            score += evaluate_window(window, piece)

    for c in range(COLS):
        col_array = [int(i) for i in list(board[:,c])]
        for r in range(ROWS-3):
            window = col_array[r:r+4]
            score += evaluate_window(window, piece)

    for r in range(3,ROWS):
        for c in range(COLS - 3):
            window = [board[r-i][c+i] for i in range(4)]
            score += evaluate_window(window, piece)

    for r in range(3,ROWS):
        for c in range(3,COLS):
            window = [board[r-i][c-i] for i in range(4)]
            score += evaluate_window(window, piece)

    return score


def is_terminal_node(board):
    return winning_move(board, PLAYER_PIECE) or winning_move(board, AI_PIECE) or len(get_valid_locations(board)) == 0


def minimax(board, depth, alpha, beta, maximizing_player):
    valid_locations = get_valid_locations(board)
    is_terminal = is_terminal_node(board)
    if depth == 0 or is_terminal:
        if is_terminal: 
            if winning_move(board, AI_PIECE):
                return (None, 10000000)
            elif winning_move(board, PLAYER_PIECE):
                return (None, -10000000)
            else:
                return (None, 0)
        else: 
            return (None, score_position(board, AI_PIECE))

    if maximizing_player:
        value = -math.inf
        column = random.choice(valid_locations)

        for col in valid_locations:
            row = get_next_open_row(board, col)
            b_copy = board.copy()
            drop_piece(b_copy, row, col, AI_PIECE)
            new_score = minimax(b_copy, depth-1, alpha, beta, False)[1]
            if new_score > value:
                value = new_score
                column = col
            alpha = max(value, alpha) 
            if alpha >= beta:
                break

        return column, value

    else:
        value = math.inf
        column = random.choice(valid_locations)
        for col in valid_locations:
            row = get_next_open_row(board, col)
            b_copy = board.copy()
            drop_piece(b_copy, row, col, PLAYER_PIECE)
            new_score = minimax(b_copy, depth-1, alpha, beta, True)[1]
            if new_score < value:
                value = new_score
                column = col
            beta = min(value, beta) 
            if alpha >= beta:
                break
        return column, value

def get_valid_locations(board):
    valid_locations = []
    
    for column in range(COLS):
        if is_valid_location(board, column):
            valid_locations.append(column)

    return valid_locations

def end_game():
    global game_over
    game_over = True

def init():
    global ROWS,COLS,PLAYER_TURN,AI_PIECE,AI_TURN,PLAYER_PIECE,board,game_over,not_over,turn
    ROWS = 7
    COLS = 7
    PLAYER_TURN = 0
    AI_TURN = 1
    PLAYER_PIECE = 0
    AI_PIECE = 2

    board = create_board()
    game_over = False
    not_over = True
    turn = 1
    draw_board(board)
    pygame.display.update()

def savePlay(board,col):
    f = open("python/data.csv","a")
    t = ""
    for x in range(7):
        for y in range(7):
            t += str(int(board[x,y])) + ","
    t += str(col) + "\n"
    f.write(t)
    f.close()


board = create_board()
game_over = False
not_over = True
turn = 1
pygame.init()
SQUARESIZE = 100
width = COLS * SQUARESIZE
height = (ROWS + 1) * SQUARESIZE
circle_radius = int(SQUARESIZE/2 - 5)
size = (width, height)
screen = pygame.display.set_mode(size)
my_font = pygame.font.SysFont("monospace", 75)
draw_board(board)
pygame.display.update()
init()

while True:
    while not game_over:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                sys.exit()
            if not_over:
                pygame.draw.rect(screen, BLACK, (0,0, width, SQUARESIZE))

        if turn == PLAYER_TURN and not_over:
            if random.randint(0,100)<20:
                col = random.randint(0,6)
                while not is_valid_location(board, col):
                    col = random.randint(0,6)
                row = get_next_open_row(board, col)
                drop_piece(board, row, col, PLAYER_PIECE)
                if winning_move(board, PLAYER_PIECE):
                    print("PLAYER 1 WINS!")
                    not_over = False
                    end_game()
            else:
                col, minimax_score = minimax(board, 5, -math.inf, math.inf, False)
                if is_valid_location(board, col):
                    savePlay(board,col)
                    row = get_next_open_row(board, col)
                    drop_piece(board, row, col, PLAYER_PIECE)
                    if winning_move(board, PLAYER_PIECE):
                        print("PLAYER 1 WINS!")
                        not_over = False
                        end_game()
            draw_board(board)
            turn += 1
            turn = turn % 2

        if turn == AI_TURN and not game_over and not_over:
            if random.randint(0,100)<20:
                col = random.randint(0,6)
                while not is_valid_location(board, col):
                    col = random.randint(0,6)
                row = get_next_open_row(board, col)
                drop_piece(board, row, col, AI_PIECE)
                if winning_move(board, AI_PIECE):
                    print("PLAYER 2 WINS!")
                    not_over = False
                    end_game()
            else:
                col, minimax_score = minimax(board, 5, -math.inf, math.inf, True)
                if is_valid_location(board, col):
                    savePlay(board,col)
                    row = get_next_open_row(board, col)
                    drop_piece(board, row, col, AI_PIECE)
                    if winning_move(board, AI_PIECE):
                        print("PLAYER 2 WINS!")
                        not_over = False
                        end_game()
            draw_board(board)
            turn += 1
            turn = turn % 2

        pygame.display.update()
    init()

