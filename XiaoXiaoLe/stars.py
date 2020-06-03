import pygame
import random
import sys
import os

# 设置页面宽高
SCRRR_WIDTH = 900
SCRRR_HEIGHT = 900
# 创建控制游戏结束的状态
GAMEOVER = False

ROW_COUNT = 20
COLUMN_COUNT = 20
COLORS = ["red","blue","green"]

LOG = '文件:{}中的方法:{}出错'.format(__file__, __name__)

def resource_path(relative_path):
    # pyinsatller静态资源打包后，获取图片路径
    if getattr(sys, 'frozen', False): #是否Bundle Resource
        base_path = sys._MEIPASS
    else:
        base_path = os.path.abspath(".")
    return os.path.join(base_path, relative_path)

class Star(pygame.sprite.Sprite):

    def __init__(self, color, location):
        super(Star, self).__init__()
        self.color = color
        self.location = location
        self.live = True
        self.image = pygame.image.load(resource_path(os.path.join("images",f"{color}.png")))
        self.rect = self.image.get_rect()
        self.rect.y = location[0] * SCRRR_HEIGHT // ROW_COUNT
        self.rect.x = location[1] * SCRRR_WIDTH // COLUMN_COUNT

    def load_image(self):
        if hasattr(self, 'image') and hasattr(self, 'rect'):
            MainGame.window.blit(self.image, self.rect)
        else:
            print(LOG)

    def change_rect(self,location):
        self.location = location
        self.rect.x = location[1] * SCRRR_WIDTH // COLUMN_COUNT
        self.rect.y = location[0] * SCRRR_HEIGHT // ROW_COUNT

    def remove(self, stars_list,click=False):
        if not self.live:
            return
        self.live = False
        row, column = self.location
        left_star = None
        right_star = None
        top_star = None
        bottom_star = None
        if column > 0:
            left_star = stars_list[row][column-1]
        if column < COLUMN_COUNT - 1:
            right_star = stars_list[row][column+1]
        if row > 0:
            top_star = stars_list[row-1][column]
        if row < ROW_COUNT - 1:
            bottom_star = stars_list[row+1][column]
        friends = 0
        for star in [left_star, right_star, top_star, bottom_star]:
            if star and star.live and star.color == self.color:
                friends += 1
                star.remove(stars_list)
        if friends == 0 and click:
            self.live = True

class MainGame:
    stars_list = []

    def start_game(self):
        self.init_window()
        self.init_stars()

        while not GAMEOVER:
            MainGame.window.fill((255, 255, 255))
            self.load_stars()
            self.deal_events()
            pygame.time.wait(20)
            pygame.display.update()

    def init_window(self):
        pygame.display.init()
        MainGame.window = pygame.display.set_mode([SCRRR_WIDTH, SCRRR_HEIGHT])

    def load_stars(self):
        for row in MainGame.stars_list:
            for star in row:
                if star.live:
                    star.load_image()
                else:
                    pass

    def init_stars(self):
        for row in range(ROW_COUNT):
            self.stars_list.append([])
            for column in range(COLUMN_COUNT):
                color = random.choice(COLORS)
                location = [row,column]
                self.stars_list[row].append(Star(color, location))

    def deal_events(self):
        eventlist = pygame.event.get()
        for e in eventlist:
            if e.type == pygame.QUIT:
                # 退出按钮
                global GAMEOVER
                GAMEOVER = True
                return self.over('游戏结束')
            elif e.type == pygame.MOUSEBUTTONDOWN:
                # 点击小星星
                if e.button == 1:
                    row = e.pos[1] // (SCRRR_HEIGHT // ROW_COUNT)
                    column = e.pos[0] // (SCRRR_WIDTH // COLUMN_COUNT)
                    self.stars_list[row][column].remove(self.stars_list,click=True)
                self.refresh_stars()

    def refresh_stars(self):
        lives_count = 0
        # null_rows_index = []
        tmp_stars_list = self.stars_list.copy()
        for row_t in tmp_stars_list:
            row_tmp = row_t.copy()
            row_lives_count = 0
            for star in row_tmp:
                if not star.live:
                    row_t.remove(star)
                    row_t.append(star)
                else:
                    lives_count += 1
                    row_lives_count += 1
            if row_lives_count == 0:
                self.stars_list.remove(row_t)
                self.stars_list.append(row_t)


        # for index in range(len(self.stars_list)):
        #     row_tmp = self.stars_list[index].copy()
        #     row_lives_count = 0
        #     for star in row_tmp:
        #         if not star.live:
        #             self.stars_list[index].remove(star)
        #             self.stars_list[index].append(star)
        #         else:
        #             lives_count += 1
        #             row_lives_count += 1
        #     if row_lives_count == 0:
        #         self.stars_list.append(self.stars_list[index].copy())
        #         null_rows_index.append(index)
        # for index in null_rows_index:
        #     del self.stars_list[index]

        if lives_count == 0:
            global GAMEOVER
            GAMEOVER = True
            return self.over('你赢了',wiat_time=800)

        for row in range(ROW_COUNT):
            for column in range(COLUMN_COUNT):
                self.stars_list[row][column].change_rect([row,column])

    # 文本绘制
    def draw_text(self, content, size, color):
        pygame.font.init()
        font = pygame.font.SysFont('kaiti', size)
        text = font.render(content, True, color)
        return text

    def over(self,msg,wiat_time=200):
        MainGame.window.fill((255, 255, 255))
        MainGame.window.blit(self.draw_text(msg, 50, (255, 0, 0)), (380, 200))
        pygame.display.update()
        pygame.time.wait(wiat_time)

if __name__ == '__main__':
    MainGame().start_game()