from __future__ import print_function

from django.shortcuts import render_to_response, HttpResponse

from host_manage import forms


# Create your views here.


def register(request):
    register_form = forms.Register()
    if request.method == 'POST':
        form = forms.Register(request.POST)
        if form.is_valid():
            data = form.cleaned_data
            user = data['username']
            pwd = data['password']
            email = data['eamil']
            print(data)
            return HttpResponse("register success!")
        else:
            print(type(form.errors))
            return HttpResponse("register failed!")
        django.forms.utils.ErrorDict

    else:
        return render_to_response('register.html', {'forms': register_form})


def login(request):
    login_form = forms.Login()
    if request.method == 'POST':
        form = forms.Login(request.POST)
        if form.is_valid():
            data = form.cleaned_data
            print(data)
            return HttpResponse('login success!')
        else:
            print (form.errors)
    else:
        return render_to_response('login.html', {'forms': login_form})
