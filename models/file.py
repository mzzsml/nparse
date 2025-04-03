import os
import pathlib

class File:
    def __init__(self, path=None):
        self.path = path
        self.content = self.getcontent()
        self.size = self.getsize()
    
    def exists(self) -> bool:
        # If file exists and it's a file, return True, else False.
        #return True if pathlib.Path(self.path).is_file() else False
        try:
            return pathlib.Path(self.path).is_file()
        except FileNotFoundError:
            print(f'error: {self.path} not found.')

    def getcontent(self):
        if self.exists():
            with open(file=self.path, mode='r') as f:
                return f.read()

    def getsize(self):
        try:
            return os.stat(self.path).st_size
        except FileNotFoundError:
            print(f'error: {self.path} does not exists.')
    
    def deletecontent(self):
        # Opening a file in write mode will delete it's content.
        if self.exists():
            open(self.path, 'w').close()
