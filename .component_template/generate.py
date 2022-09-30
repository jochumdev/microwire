#!./venv/bin/python3

import os
import os.path
import click
import glob
from jinja2 import Environment, FileSystemLoader, select_autoescape
from yaml import load as YamlLoad
try:
    from yaml import CLoader as YamlLoader
except ImportError:
    from yaml import YamlLoader

@click.command()
@click.option('--out_dir', help="Output directory, by default $script_path/dist/")
@click.argument('config', type=click.Path(exists=True))
def generate(config, out_dir):
    script_path = os.path.dirname(os.path.abspath(__file__))
    render_dir = os.path.join(script_path, 'dist')
    env = Environment(
        loader=FileSystemLoader("templates"),
        autoescape=select_autoescape()
    )

    with open(config, 'r+b') as fp:
        data = YamlLoad(fp, YamlLoader)
        
        out_path = os.path.join(render_dir, data['Name'])

        if out_dir != "":
            out_path = out_dir

        os.makedirs(out_path, exist_ok=True)

        templates = glob.glob('*.j2', root_dir=os.path.join(script_path, 'templates'))
        for template in templates:
            with open(os.path.join(out_path, os.path.basename(template)[:-3]), 'w') as wp:
                out_text = env.get_template(os.path.basename(template)).render(**data)
                out_text += "\n"
                wp.write(out_text)


if __name__ == '__main__':
    generate()
