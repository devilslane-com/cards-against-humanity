<?php

namespace CAH\Models;

use MongoDB\Laravel\Relations;

class Round extends BaseModel
{
    protected $casts = [
        'started_at' => 'datetime',
        'ended_at' => 'datetime',
    ];

    public function challenge () : Relations\HasOne 
    {
        return $this->hasOne (Card::class, 'challenge');
    }

    public function dealer () : Relations\HasOne 
    {
        return $this->hasOne (Player::class, 'dealer');
    }

    public function game () : Relations\BelongsTo 
    {
        return $this->belongsTo (Game::class);
    }

    public function responses () : Relations\HasMany
    {
        return $this->hasMany (Response::class);
    }

    public function winner () : Relations\HasOne 
    {
        return $this->hasOne (Player::class, 'winner');
    }
}
