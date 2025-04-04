import os
import json
import pathlib

import lib
import lib.tojsonl

class File:
    """File object that handles operations with files."""

    def __init__(self, path=None):
        self.path = path
        self.content = self.getcontent()
        self.size = self.getsize()
    
    def exists(self) -> bool:
        """Check wether a file exists."""

        # Because during the creation of the output file we append each json
        # line, we need to implement a basic function that checks wether the file exists,
        # and and returns True if it does or False if it doesn't.
        return True if pathlib.Path(self.path).is_file() else False

    def getcontent(self):
        """Get the file content."""

        if self.exists():
            with open(file=self.path, mode='r') as f:
                return f.read()

    def getsize(self):
        """Get the file size."""

        if self.exists():
            return os.stat(self.path).st_size
        else:
            return None
    
    def deletecontent(self):
        """Delete the file content."""

        # Opening a file in write mode will delete it's content.
        if self.exists():
            open(self.path, 'w').close()

    def writecontent(self, content):
        """Write content into a file."""

        with open(self.path, '+a') as f:
            [json.dump(x, f) for x in lib.tojsonl.tojsonl(content)]