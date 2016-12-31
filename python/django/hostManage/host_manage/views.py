from __future__ import print_function

from django.shortcuts import render_to_response, HttpResponse, redirect

from host_manage import forms, models


# Create your views here.


def register(request):
    register_form = forms.Register()
    if request.method == 'POST':
        role = request.POST.get('role', '')
        form = forms.Register(request.POST)
        if form.is_valid():
            data = form.cleaned_data
            user = data.get('username', '')
            pwd = data.get('password', '')
            email = data.get('email', '')
            group_obj = models.UserGroup.objects.get(name=role)
            info_obj = models.UserInfo.objects.create(name=user, password=pwd, email=email)
            info_obj.userGroup.add(group_obj)
            message = "register success!" if info_obj else "register failed!"
            return HttpResponse(message)
        else:
            print(form.errors)
            return HttpResponse("register failed!")

    else:
        return render_to_response('register.html', {'forms': register_form})


def login(request):
    login_form = forms.Login()
    message = ''
    if request.method == 'POST':
        form = forms.Login(request.POST)
        if form.is_valid():
            data = form.cleaned_data
            user = data.get('username', '')
            pwd = data.get('password', '')
            user_group = models.UserInfo.objects.filter(name=user, password=pwd)
            # print(user_group)
            if user_group:
                request.session['user'] = user
                return redirect('/hosts')
            message = 'login failed! '
        else:
            print(form.errors)
    return render_to_response('login.html', {'forms': login_form, 'message': message})


def create_host(request):
    message = ''
    create_host_forms = forms.CreateHost()
    if request.method == 'POST':
        form = forms.CreateHost(request.POST)
        if form.is_valid():
            data = form.cleaned_data
            hostname = data.get('hostname', '')
            ip = data.get('ip', '')
            user_group = data.get('host_group', 'user')
            user_group_id = models.UserGroup.objects.get(name=user_group).id
            res = models.Hosts.objects.create(hostname=hostname, ip=ip, userGroup_id=user_group_id)
            if res:
                message = "create host success!"
        else:
            message = form.errors

    return render_to_response('create_host.html', {'forms': create_host_forms, 'message': message})


def show_hosts(request):
    hosts = []
    if request.method == 'GET':
        user = request.session.get("user", "")
        if user:
            user_group_id = models.UserInfo.objects.get(name=user).userGroup.values().first()
            if user_group_id:
                hosts = models.Hosts.objects.filter(userGroup_id=user_group_id.get('id', ''))

    return render_to_response('hosts.html', {'hosts': hosts})
