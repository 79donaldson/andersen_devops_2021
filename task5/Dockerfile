FROM  python:3-alpine
ADD my_py.py /

RUN pip install --upgrade pip

RUN pip install Flask

CMD [ "python", "/my_py.py" ]

EXPOSE 5000