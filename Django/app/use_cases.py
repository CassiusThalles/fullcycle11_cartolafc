from typing import List
from app.models import MyTeam, Player

class UpdatePlayersInMyTeam:
    def execute(self, my_team_id, players_uuid: List[str]):
        my_team = MyTeam.objects.get(id=my_team_id)
        players = Player.objects.filter(uuid__in=players_uuid)
        my_team.players.set(players)