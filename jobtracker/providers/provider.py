from abc import ABC, abstractmethod
class Provider(ABC):
    @staticmethod
    @abstractmethod  
    def get_jobs():
        pass