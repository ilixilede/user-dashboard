import dash
import dash_core_components as dcc
import dash_html_components as html
from dash.dependencies import Input, Output
import plotly.express as px
import pandas as pd

app = dash.Dash(__name__)

app.layout = html.Div([
    html.H1('User Dashboard'),
    dcc.Dropdown(
        id='dropdown',
        options=[
            {'label': 'Option 1', 'value': 'option1'},
            {'label': 'Option 2', 'value': 'option2'}
        ],
        value='option1'
    ),
    dcc.Graph(id='graph')
])

@app.callback(
    Output('graph', 'figure'),
    [Input('dropdown', 'value')]
)
def update_graph(selected_value):
    df = pd.DataFrame({
        'x': ['apple', 'banana', 'cherry'],
        'y': [1, 2, 3]
    })

    fig = px.bar(df, x='x', y='y')
    return fig

if __name__ == '__main__':
    app.run_server(debug=True)