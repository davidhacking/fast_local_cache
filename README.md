# fast_local_cache
![img](./doc/fast_local_cache.png)
## 使用场景
- 推荐打分时需要item特征数据，为了实时性，会将item特征缓存起来，运算的时候通过本地计算即可
## 目标
- 全缓存保证读性能

## 数据结构
- cache key->value
- datafiles 
  - 版本号：递增，merge后例如10个datafile merge成了5个，则这5个datafile的版本都为10
  - entry ksz key del_flag vsz value ts crc

## 接口
- SyncDB同步最新数据
  - 输入：远端版本号
  - 输出：server版本号 datafiles

## 问题
- 如何进行merge？
  - merge选取一部分file进行merge，此时读取内存中的key，比较ts，如果大于则更新内存中的key，并且写磁盘新entry，如果等于则只写磁盘新entry，如果小于则直接丢弃
  - 多久触发一次？几个datafile进行merge？可配
  - merge需要加全局锁？
  - 如果merge后的版本是最新的版本，那远端每次更新最新的版本时会导致更新多个文件，所以远端需要更新轮询时间需要小于merge时间
- datafile如何load？
  - load data file mmap
- 如何进行get？
  - 读内存返回
- 如何进行put？
  - 追加式写磁盘，再写内存
- delete？和put类似

## 相关链接
- [bitcask介绍](https://doc.weixin.qq.com/slide/p3_m_PGlWCBdxkkzO?scode=AJEAIQdfAAoAjLccpeAKIAiAYdACo)