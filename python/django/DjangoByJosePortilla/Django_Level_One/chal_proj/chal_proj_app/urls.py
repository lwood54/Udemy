from django.conf.urls import url
from chal_proj_app import views

urlpatterns = [
    url(r'^$',views.index,name='index'),
    url(r'^index/$',views.index,name='index'),
    url(r'^help/$',views.help,name='help'),
]
