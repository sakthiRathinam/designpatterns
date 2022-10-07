class Os:
    def privacy(self):
        print("imhere")

class Test:
    def test_func(self):
        print("original func")
class Testt:
    def test_funcc(self):
        print("test func")


class Dog(Os, Test):
    animal = "Dog"

    def __init__(self):
        self.animal = "scream"

    def test_func(self):
        super(Dog, self).test_func()
        print("serewr")
        
print(Dog().test_func(), "here")
print(Dog.animal,"class variable")





class Parent(object):
    def __init__(self):
        super(Parent, self).__init__()
        print("parent")
class Ch(object):
    def __init__(self):
        super(Ch, self).__init__()
        print("Ch")

    def test_func(self):
        print("im in Ch")
        super().test_func()

class Left(Ch):
    def __init__(self):
        super(Left, self).__init__()

        print("left")
        
    


class Right(Parent, Ch):
    def __init__(self):
        super(Right, self).__init__()
        print("right")

    def test_func(self):
        print("im in right")

class Child(Left, Right):
    def __init__(self):
        super(Right, self).__init__()
        print("child")
    #     self.test_func()
    # def test_func(self):
    #     print("im in child")
    #     super().test_func()

c = Child()
print(Child.__mro__)
