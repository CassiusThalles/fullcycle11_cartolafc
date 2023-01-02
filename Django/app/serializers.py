from rest_framework import serializers
from .models import Player

class UpdateMyPlayerSerializer(serializers.Serializer):
    players_id = serializers.PrimaryKeyRelatedField(
        many=True,
        queryset=Player.objects.all()
    )