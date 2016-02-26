<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: query.proto
//   Date: 2016-01-22 01:34:42

namespace Vitess\Proto\Query {

  class StreamExecuteRequest extends \DrSlump\Protobuf\Message {

    /**  @var \Vitess\Proto\Vtrpc\CallerID */
    public $effective_caller_id = null;
    
    /**  @var \Vitess\Proto\Query\VTGateCallerID */
    public $immediate_caller_id = null;
    
    /**  @var \Vitess\Proto\Query\Target */
    public $target = null;
    
    /**  @var \Vitess\Proto\Query\BoundQuery */
    public $query = null;
    
    /**  @var int */
    public $session_id = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'query.StreamExecuteRequest');

      // OPTIONAL MESSAGE effective_caller_id = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "effective_caller_id";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $f->reference = '\Vitess\Proto\Vtrpc\CallerID';
      $descriptor->addField($f);

      // OPTIONAL MESSAGE immediate_caller_id = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "immediate_caller_id";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $f->reference = '\Vitess\Proto\Query\VTGateCallerID';
      $descriptor->addField($f);

      // OPTIONAL MESSAGE target = 3
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 3;
      $f->name      = "target";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $f->reference = '\Vitess\Proto\Query\Target';
      $descriptor->addField($f);

      // OPTIONAL MESSAGE query = 4
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 4;
      $f->name      = "query";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $f->reference = '\Vitess\Proto\Query\BoundQuery';
      $descriptor->addField($f);

      // OPTIONAL INT64 session_id = 5
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 5;
      $f->name      = "session_id";
      $f->type      = \DrSlump\Protobuf::TYPE_INT64;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <effective_caller_id> has a value
     *
     * @return boolean
     */
    public function hasEffectiveCallerId(){
      return $this->_has(1);
    }
    
    /**
     * Clear <effective_caller_id> value
     *
     * @return \Vitess\Proto\Query\StreamExecuteRequest
     */
    public function clearEffectiveCallerId(){
      return $this->_clear(1);
    }
    
    /**
     * Get <effective_caller_id> value
     *
     * @return \Vitess\Proto\Vtrpc\CallerID
     */
    public function getEffectiveCallerId(){
      return $this->_get(1);
    }
    
    /**
     * Set <effective_caller_id> value
     *
     * @param \Vitess\Proto\Vtrpc\CallerID $value
     * @return \Vitess\Proto\Query\StreamExecuteRequest
     */
    public function setEffectiveCallerId(\Vitess\Proto\Vtrpc\CallerID $value){
      return $this->_set(1, $value);
    }
    
    /**
     * Check if <immediate_caller_id> has a value
     *
     * @return boolean
     */
    public function hasImmediateCallerId(){
      return $this->_has(2);
    }
    
    /**
     * Clear <immediate_caller_id> value
     *
     * @return \Vitess\Proto\Query\StreamExecuteRequest
     */
    public function clearImmediateCallerId(){
      return $this->_clear(2);
    }
    
    /**
     * Get <immediate_caller_id> value
     *
     * @return \Vitess\Proto\Query\VTGateCallerID
     */
    public function getImmediateCallerId(){
      return $this->_get(2);
    }
    
    /**
     * Set <immediate_caller_id> value
     *
     * @param \Vitess\Proto\Query\VTGateCallerID $value
     * @return \Vitess\Proto\Query\StreamExecuteRequest
     */
    public function setImmediateCallerId(\Vitess\Proto\Query\VTGateCallerID $value){
      return $this->_set(2, $value);
    }
    
    /**
     * Check if <target> has a value
     *
     * @return boolean
     */
    public function hasTarget(){
      return $this->_has(3);
    }
    
    /**
     * Clear <target> value
     *
     * @return \Vitess\Proto\Query\StreamExecuteRequest
     */
    public function clearTarget(){
      return $this->_clear(3);
    }
    
    /**
     * Get <target> value
     *
     * @return \Vitess\Proto\Query\Target
     */
    public function getTarget(){
      return $this->_get(3);
    }
    
    /**
     * Set <target> value
     *
     * @param \Vitess\Proto\Query\Target $value
     * @return \Vitess\Proto\Query\StreamExecuteRequest
     */
    public function setTarget(\Vitess\Proto\Query\Target $value){
      return $this->_set(3, $value);
    }
    
    /**
     * Check if <query> has a value
     *
     * @return boolean
     */
    public function hasQuery(){
      return $this->_has(4);
    }
    
    /**
     * Clear <query> value
     *
     * @return \Vitess\Proto\Query\StreamExecuteRequest
     */
    public function clearQuery(){
      return $this->_clear(4);
    }
    
    /**
     * Get <query> value
     *
     * @return \Vitess\Proto\Query\BoundQuery
     */
    public function getQuery(){
      return $this->_get(4);
    }
    
    /**
     * Set <query> value
     *
     * @param \Vitess\Proto\Query\BoundQuery $value
     * @return \Vitess\Proto\Query\StreamExecuteRequest
     */
    public function setQuery(\Vitess\Proto\Query\BoundQuery $value){
      return $this->_set(4, $value);
    }
    
    /**
     * Check if <session_id> has a value
     *
     * @return boolean
     */
    public function hasSessionId(){
      return $this->_has(5);
    }
    
    /**
     * Clear <session_id> value
     *
     * @return \Vitess\Proto\Query\StreamExecuteRequest
     */
    public function clearSessionId(){
      return $this->_clear(5);
    }
    
    /**
     * Get <session_id> value
     *
     * @return int
     */
    public function getSessionId(){
      return $this->_get(5);
    }
    
    /**
     * Set <session_id> value
     *
     * @param int $value
     * @return \Vitess\Proto\Query\StreamExecuteRequest
     */
    public function setSessionId( $value){
      return $this->_set(5, $value);
    }
  }
}

