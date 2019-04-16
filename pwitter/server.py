from flask import Flask, jsonify, request
from pwitter import config
from pwitter.model import Pweet
from schematics.exceptions import ValidationError, ConversionError

app = Flask(__name__)

@app.route('/pweets', methods=['GET'])
def get():
    lte = float(request.args.get('lte', '1.0'))
    gte = float(request.args.get('gte', '-1.0'))

    if lte <= gte:
        return jsonify(
                {'message': 'lte cannot be less than or equal to gte'}
               ), 400

    pweets = Pweet.by_polarity(low=gte, high=lte)
    body = {
        'pweets': [p.to_primitive() for p in pweets]
    }
    return jsonify(body)

@app.route('/pweets/count', methods=['GET'])
def count():
    return jsonify({'count': Pweet.count()})

@app.route('/pweets', methods=['POST'])
def post():
    '''
    This expects a Pweet and stores it to database
    '''
    pweet = Pweet(request.form)
    pweet.validate()
    pweet.save()

    return jsonify(pweet.serialize()), 201

@app.errorhandler(ValidationError)
@app.errorhandler(ConversionError)
def handle_validation_error(error):
    resp = jsonify({'message': error.message})
    resp.status_code = 400
    return resp

if __name__ == '__main__':
    app.run(
        debug=True,
        host=config.get('WEB_ADDRESS'),
        port=int(config.get('WEB_PORT'))
    )
