import logging
from logging.config import dictConfig
import os


DIRECTORY_CONN_STR = "dbname='{}' user='{}' host='{}' password='{}'".format(
    'directory',
    os.getenv('DIPLO_DIRECTORY_DB_USER'),
    os.getenv('DIPLO_DIRECTORY_DB_HOST'),
    os.getenv('DIPLO_DIRECTORY_DB_PASSWORD'))


CHAT_CONN_STR = "dbname='{}' user='{}' host='{}' password='{}'".format(
    'chat',
    os.getenv('DIPLO_CHAT_DB_USER'),
    os.getenv('DIPLO_CHAT_DB_HOST'),
    os.getenv('DIPLO_CHAT_DB_PASSWORD'))


DIRECTORY_TABLES = [
    'account'
]


CHAT_TABLES = [
    'message',
    'channel_member',
    'channel',
    'chat_user'
]


BASE_HEADERS = {
    'Content-Type': 'application/json'
}


USER_1 = {
    "email": "metternic@gmail.com",
    "username": "mr.austria",
    "password": "password",
    "givenName": "Clemens",
    "surname": "Von Metternic"
}


USER_2 = {
    "email": "castelraegh@gmail.com",
    "username": "mr.england",
    "password": "password",
    "givenName": "Robert Stewart",
    "surname": "Castelraegh"
}


CHANNEL_NAME = 'The Melian dialoge'


BASE_URL = 'http://localhost:8080/api'


# Directory routes
CREATE_USER_ROUTE = f'{BASE_URL}/directory/user'
LOGIN_USER_ROUTE = f'{BASE_URL}/directory/user/login'
RENEW_TOKEN_ROUTE = f'{BASE_URL}/directory/user/token/renew'
GET_USER_ROUTE = f'{BASE_URL}/directory/me'
GET_USERS_ROUTE = f'{BASE_URL}/directory/users'

# Chat routes
CHAT_USER_ROUTE = f'{BASE_URL}/chat/user'
CREATE_CHANNEL_ROUTE = f'{BASE_URL}/chat/channel/' '{}' f'/{CHANNEL_NAME}'
CHANNEL_ROUTE = f'{BASE_URL}/chat/channel/' '{}'
GET_CHANNEL_ROUTE = f'{BASE_URL}/chat/channel/' '{}'
MESSAGE_ROUTE = f'{BASE_URL}/chat/message/' '{}'


LOGGING_CONIFG = {
    'version': 1,
    'formatters': {
        'default': {
            'format': '%(asctime)s - %(name)s - %(levelname)s - %(message)s'
        }
    },
    'handlers': {
        'console': {
            'class': 'logging.StreamHandler',
            'formatter': 'default',
            'level': logging.INFO
        }
    },
    'root': {
        'handlers': ['console'],
        'level': logging.DEBUG
    }
}


dictConfig(LOGGING_CONIFG)


MELIAN_DIALOGE = [
    'Since the negotiations are not to go on before the people, in order that we may not be able to speak straight on without interruption, and deceive the ears of the multitude by seductive arguments which would pass without refutation (for we know that this is the meaning of our being brought before the few), what if you who sit there were to pursue a method more cautious still? Make no set speech yourselves, but take us up at whatever you do not like, and settle that before going any farther. And first tell us if this proposition of ours suits you.',
    'To the fairness of quietly instructing each other as you propose there is nothing to object; but your military preparations are too far advanced to agree with what you say, as we see you are come to be judges in your own cause, and that all we can reasonably expect from this negotiation is war, if we prove to have right on our side and refuse to submit, and in the contrary case, slavery.',
    'If you have met to reason about presentiments of the future, or for anything else than to consult for the safety of your state upon the facts that you see before you, we will give over; otherwise we will go on.',
    'It is natural and excusable for men in our position to turn more ways than one both in thought and utterance. However, the question in this conference is, as you say, the safety of our country; and the discussion, if you please, can proceed in the way which you propose.',
    'For ourselves, we shall not trouble you with specious pretences- either of how we have a right to our empire because we overthrew the Mede, or are now attacking you because of wrong that you have done us- and make a long speech which would not be believed; and in return we hope that you, instead of thinking to influence us by saying that you did not join the Lacedaemonians, although their colonists, or that you have done us no wrong, will aim at what is feasible, holding in view the real sentiments of us both; since you know as well as we do that right, as the world goes, is only in question between equals in power, while the strong do what they can and the weak suffer what they must.',
    'As we think, at any rate, it is expedient- we speak as we are obliged, since you enjoin us to let right alone and talk only of interest- that you should not destroy what is our common protection, the privilege of being allowed in danger to invoke what is fair and right, and even to profit by arguments not strictly valid if they can be got to pass current. And you are as much interested in this as any, as your fall would be a signal for the heaviest vengeance and an example for the world to meditate upon.',
    'The end of our empire, if end it should, does not frighten us: a rival empire like Lacedaemon, even if Lacedaemon was our real antagonist, is not so terrible to the vanquished as subjects who by themselves attack and overpower their rulers. This, however, is a risk that we are content to take. We will now proceed to show you that we are come here in the interest of our empire, and that we shall say what we are now going to say, for the preservation of your country; as we would fain exercise that empire over you withouttrouble, and see you preserved for the good of us both.',
    'And how, pray, could it turn out as good for us to serve as for you to rule?',
    'Because you would have the advantage of submitting before suffering the worst, and we should gain by not destroying you.',
    'So that you would not consent to our being neutral, friends instead of enemies, but allies of neither side.',
    'No; for your hostility cannot so much hurt us as your friendship will be an argument to our subjects of our weakness, and your enmity of our power.',
    "Is that your subjects' idea of equity, to put those who have nothing to do with you in the same category with peoples that are most of them your own colonists, and some conquered rebels?",
    'As far as right goes they think one has as much of it as the other, and that if any maintain their independence it is because they are strong, and that if we do not molest them it is because we are afraid; so that besides extending our empire we should gain in security by your subjection; the fact that you are islanders and weaker than others rendering it all the more important that you should not succeed in baffling the masters of the sea.',
    'But do you consider that there is no security in the policy which we indicate? For here again if you debar us from talking about justice and invite us to obey your interest, we also must explain ours, and try to persuade you, if the two happen to coincide. How can you avoid making enemies of all existing neutrals who shall look at case from it that one day or another you will attack them? And what is this but to make greater the enemies that you have already, and to force others to become so who would otherwise have never thought of it?',
    'Why, the fact is that continentals generally give us but little alarm; the liberty which they enjoy will long prevent their taking precautions against us; it is rather islanders like yourselves, outside our empire, and subjects smarting under the yoke, who would be the most likely to take a rash step and lead themselves and us into obvious danger.',
    'Well then, if you risk so much to retain your empire, and your subjects to get rid of it, it were surely great baseness and cowardice in us who are still free not to try everything that can be tried, before submitting to your yoke.',
    'Not if you are well advised, the contest not being an equal one, with honour as the prize and shame as the penalty, but a question of self-preservation and of not resisting those who are far stronger than you are.',
    'But we know that the fortune of war is sometimes more impartial than the disproportion of numbers might lead one to suppose; to submit is to give ourselves over to despair, while action still preserves for us a hope that we may stand erect.',
    "Hope, danger's comforter, may be indulged in by those who have abundant resources, if not without loss at all events without ruin; but its nature is to be extravagant, and those who go so far as to put their all upon the venture see it in its true colours only when they are ruined; but so long as the discovery would enable them to guard against it, it is never found wanting. Let not this be the case with you, who are weak and hang on a single turn of the scale; nor be like the vulgar, who, abandoning such security as human means may still afford, when visible hopes fail them in extremity, turn to invisible, to prophecies and oracles, and other such inventions thatdelude men with hopes to their destruction.",
    'You may be sure that we are as well aware as you of the difficulty of contending against your power and fortune, unless the terms be equal. But we trust that the gods may grant us fortune as good as yours, since we are just men fighting against unjust, and that what we want in power will be made up by the alliance of the Lacedaemonians, who are bound, if only for very shame, to come to the aid of their kindred. Our confidence, therefore, after all is not so utterly irrational.',
    "When you speak of the favour of the gods, we may as fairly hope for that as yourselves; neither our pretensions nor our conduct being in any way contrary to what men believe of the gods, or practise among themselves. Of the gods we believe, and of men we know, that by a necessary law of their nature they rule wherever they can. And it is not as if we were the first to make this law, or to act upon it when made: we found it existing before us, and shall leave it to exist for ever after us; all we do is to make use of it, knowing that you and everybody else, having the same power as we have, would do the same as we do. Thus, as far as the gods are concerned, we have no fear and no reason to fear that we shall be at a disadvantage. But when we come to your notion about the Lacedaemonians, which leads you to believe that shame will make them help you, here we bless your simplicity but do not envy your folly. The Lacedaemonians, when their own interests or their country's laws are in question, are the worthiest men alive; of their conduct towards others much might be said, but no clearer idea of it could be given than by shortly saying that of all the men we know they are most conspicuous in considering what is agreeable honourable, and what is expedient just. Such a way of thinking does not promise much for the safety which you now unreasonably count upon.",
    'But it is for this very reason that we now trust to their respect for expediency to prevent them from betraying the Melians, their colonists, and thereby losing the confidence of their friends in Hellas and helping their enemies.',
    'Then you do not adopt the view that expediency goes with security, while justice and honour cannot be followed without danger; and danger the Lacedaemonians generally court as little as possible.',
    'But we believe that they would be more likely to face even danger for our sake, and with more confidence than for others, as our nearness to Peloponnese makes it easier for them to act, and our common blood ensures our fidelity.',
    'Yes, but what an intending ally trusts to is not the goodwill of those who ask his aid, but a decided superiority of power for action; and the Lacedaemonians look to this even more than others. At least, such is their distrust of their home resources that it is only with numerous allies that they attack a neighbour; now is it likely that while we are masters of the sea they will cross over to an island?',
    'But they would have others to send. The Cretan Sea is a wide one, and it is more difficult for those who command it to intercept others, than for those who wish to elude them to do so safely. And should the Lacedaemonians miscarry in this, they would fall upon your land, and upon those left of your allies whom Brasidas did not reach; and instead of places which are not yours, you will have to fight for your own country and your own confederacy.',
    'Some diversion of the kind you speak of you may one day experience, only to learn, as others have done, that the Athenians never once yet withdrew from a siege for fear of any. But we are struck by the fact that, after saying you would consult for the safety of your country, in all this discussion you have mentioned nothing which men might trust in and think to be saved by. Your strongest arguments depend upon hope and the future, and your actual resources are too scanty, as compared with those arrayed against you, for you to come out victorious. You will therefore show great blindness of judgment, unless, after allowing us to retire, you can find some counsel more prudent than this. You will surely not be caught by that idea of disgrace, which in dangers that are disgraceful, and at the same time too plain to be mistaken, proves so fatal to mankind; since in too many cases the very men that have their eyes perfectly open to what they are rushing into, let the thing called disgrace, by the mere influence of a seductive name, lead them on to a point at which they become so enslaved by the phrase as in fact to fall wilfully into hopeless disaster, and incur disgrace more disgraceful as the companion of error, than when it comes as the result of misfortune. This, if you are well advised, you will guard against; and you will not think it dishonourable to submit to the greatest city in Hellas, when it makes you the moderate offer of becoming its tributary ally, without ceasing to enjoy the country that belongs to you; nor when you have the choice given you between war and security, will you be so blinded as to choose the worse. And it is certain that those who do not yield to their equals, who keep terms with their superiors, and are moderate towards their inferiors, on the whole succeed best. Think over the matter, therefore, after our withdrawal, and reflect once and again that it is for your country that you are consulting, that you have not more than one, and that upon this one deliberation depends its prosperity or ruin.'
]
