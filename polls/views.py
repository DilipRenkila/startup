from django.http import HttpResponse

def index(request):
    return HttpResponse("this is the index page of polls app.")