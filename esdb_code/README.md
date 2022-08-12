# esdb
项目功能： 同步mysql数据到es

## 重要说明:
Elasticsearch 目前版本为 7.1.1 , 所有的插件也必须与 ES 版本相同

## 包含插件：
- 分词器
- 拼音分词器
- 繁简转换

##以下文章为使用ik中文分词

Elasticsearch 在进行存储时，会对文章内容字段进行分词，获取并保存分词后的词元（tokens）

1、安装插件
```
./bin/elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.1.1/elasticsearch-analysis-ik-7.1.1.zip


./bin/elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v6.3.0/elasticsearch-analysis-ik-6.3.0.zip

```
2、修改配置文件
在ik目录下的config目录下，IkAnalyzer.cfg.xml可以配置扩展词，及远程扩展字典。
```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE properties SYSTEM "http://java.sun.com/dtd/properties.dtd">
<properties>
        <comment>IK Analyzer 扩展配置</comment>
        <!--用户可以在这里配置自己的扩展字典 -->
        <entry key="ext_dict">ext_dict.dic</entry>
         <!--用户可以在这里配置自己的扩展停止词字典-->
        <entry key="ext_stopwords"></entry>
        <!--用户可以在这里配置远程扩展字典 -->
        <!-- <entry key="remote_ext_dict">words_location</entry> -->
        <!--用户可以在这里配置远程扩展停止词字典-->
        <!-- <entry key="remote_ext_stopwords">words_location</entry> -->
</properties>
```

扩展词库
vim ext_dict.dic （可以从搜索日志来，编码utf-8，换行使用unix格式，千万不要window的\r）
影片1
影片2

3、重启es

4、测试分词细读颗粒，ik_smart和ik_max_word
```
curl -XPOST http://127.0.0.1:9200/_analyze -d `{
    "analyzer":"ik_max_word",
    "text":"中国人民共和国"
}`

```

6，确认安装完成后，重新运行esdb程序，生成新的词元
先删除之前的索引
```
./esdb del movie
./esdb del movie_actor
./esdb del movie_category
./esdb del movie_director
./esdb del movie_film_companies
./esdb del movie_label
./esdb del movie_series
```

重新创建带新分词规则的索引
./esdb -d=true


# 以下文章为使用拼音分词器

1、安装插件
```
./bin/elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-pinyin/releases/download/v7.1.1/elasticsearch-analysis-pinyin-7.1.1.zip
```

2、测试拼音分词器
```
curl -XPOST http://127.0.0.1:9200/_analyze -d `{
      "analyzer":"pinyin",
      "text":"学习"
}`
```

3、其它步骤与分词器相同

## 繁体简体转换
1、安装插件
```
./bin/elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-stconvert/releases/download/v7.1.1/elasticsearch-analysis-stconvert-7.1.1.zip
```
2、测试分词器
```
curl -XPOST http://127.0.0.1:9200/_analyze -d `{
      "tokenizer" : "keyword",
      "char_filter" : ["stconvert"],
      "text" : "亲爱的安德烈，中国"
}`
```