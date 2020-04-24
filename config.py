# Server Related Settings

"""-----------------------------------------------------------------------------
STATUS only has one significant value: 'dev'. If STATUS is set to 'dev', Morty
starts server in development mode. Otherwise, 'waitress' production server will
be utitlized.
-----------------------------------------------------------------------------"""
STATUS: str = 'dep'

""" Server port. """
PORT: int = 5050
