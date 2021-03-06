# 概要

ビルを建てる時は、まず地盤を固め、骨組みを作り、下から上に少しずつ作っていきます。  
一般的に複雑な構造を持ったものを作る時には、段階を踏み、少しずつ完成させていきます。  
Builderパターンはそんな複雑な構造を持ったインスタンスを作成する際に使うパターンです。  

# 誰が何を知っているか

オブジェクト指向プログラミングにおいて、誰が何を知っているかはとても大切です。  

- 呼び出し側： Builderインターフェイスについて何も知らない。DirectorでBuild()を呼び出すだけ。
- Director: Builderインターフェイスで定義されているメソッドを使って文章構築する。実際にどういう文章を作成するかは知らない。
TextBuilder, HTMLBuilderがどういうものかは一切関与しない。

上記のように詳しく知らないことによって、後から入れ替えが自由に出来る。これを交換可能生といいます。

# Builderのインターフェイスを考えるポイント

Builderクラスは、文章を構築するにに、必要なメソッドをインターフェイスとして用意する必要があります。なぜなら、Directorに与えられる道具は、Builderクラスが提供するものだからです。  
将来増える可能生のあるものにBuilderは対応する必要があります。今回はシンプルな構成でしたが、無限に増える仕様に対応出来るでしょうか？しかしながら、ある程度、どういうものがありえるかを見越すことは重要です。
