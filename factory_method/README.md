# 概要

Template Mehodパターンでは、スーパクラス側で処理の骨組みを作り、サブクラス側で具体的な処理の肉付けを行いました。このパターンをインスタンス生成の場面に適用したのがFactory Methodパターンです。

Factoryという単語の意味は、工場です。インスタンスを生成する工場をTemplate Methodパターンで構成したものが、Factory Methodになります。

Factory Methodは、インスタンスの作り方をスーパークラス側で定めます。  
具体的な肉付けはサブクラスに任せています。

# framework側は依存がない

いま作成しているものとは別にテレビクラスのようなものを作った時に、framework側を一切修正せずに、全く別の製品と工場が作れることに注目してみると、どんなものが増えたとしても、frameworkは何にも依存しません。

