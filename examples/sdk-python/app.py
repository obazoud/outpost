import typer

from example.create_destination import run as create_destination
from example.auth import run as auth

app = typer.Typer()

app.command("auth")(auth)
app.command("create-destination")(create_destination)


if __name__ == "__main__":
    app()
